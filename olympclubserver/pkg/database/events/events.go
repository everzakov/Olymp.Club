package events

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EventTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type EventUserTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

var (
	ErrConnectionExists      = fmt.Errorf("Connection already exists")
	ErrConnectionDoesntExist = fmt.Errorf("Connection does not exist")
	ErrEventIsAlreadyExisted = fmt.Errorf("Event is already existed")
)

type Event struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Short       string `json:"short"`
	Img         string `json:"img"`
	Status      string `json:"status"`
	HolderID    int32  `json:"holder_id"`
	Website     string `json:"website"`
}

func (table *EventTable) CreateEvent(event Event) (Event, error) {
	filter := NewEventFilter()
	filter.Short = event.Short
	events, err := table.GetEvents(filter)
	if err != nil {
		return Event{}, err
	}
	if len(events) != 0 {
		return Event{}, ErrEventIsAlreadyExisted
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	err = table.Connection.QueryRow(context.Background(), "insert into \"EventModel\" (\"Name\", \"Description\", \"Short\", \"Img\", \"Status\", \"HolderId\", \"Website\") VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;", event.Name, event.Description, event.Short, event.Img, event.Status, event.HolderID, event.Website).Scan(&event.ID)
	// fmt.Println(err)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func (table *EventTable) GetEvents(filter EventFilter) ([]Event, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	events := []Event{}
	queryString, queryArgs := GetQueryEventOptions(filter)
	rows, err := table.Connection.Query(context.Background(), queryString, queryArgs...)
	// fmt.Println(err)
	if err != nil {
		return []Event{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		event := Event{}
		event.ID = values[0].(int32)
		event.Name = values[1].(string)
		event.Description = values[2].(string)
		event.Short = values[3].(string)
		event.Img = "img/" + values[4].(string)
		event.Status = values[5].(string)
		event.HolderID = values[6].(int32)
		event.Website = values[7].(string)
		events = append(events, event)
	}
	return events, nil
}

func (table *EventUserTable) GetEvents(userID int32) ([]Event, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select em.* from \"EventUserModel\" as eu left outer join \"EventModel\" as em on eu.event_id = em.id where eu.user_id=$1", userID)
	// fmt.Println(err)
	if err != nil {
		return []Event{}, err
	}
	events := []Event{}
	for rows.Next() {
		values, err := rows.Values()
		// fmt.Println(err)
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		// fmt.Println(len(values), values)
		event := Event{}
		event.ID = values[0].(int32)
		event.Name = values[1].(string)
		event.Description = values[2].(string)
		event.Short = values[3].(string)
		event.Img = "img/" + values[4].(string)
		event.Status = values[5].(string)
		event.HolderID = values[6].(int32)
		event.Website = values[7].(string)
		events = append(events, event)
	}
	return events, nil
}

func (table *EventUserTable) CreateConnection(userID, eventID int32) error {
	events, err := table.GetEvents(userID)
	if err != nil {
		return err
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	check := false
	for _, event := range events {
		if event.ID == eventID {
			check = true
		}
	}
	if check {
		return ErrConnectionExists
	}
	err = table.Connection.QueryRow(context.Background(), "insert into \"EventUserModel\" (event_id, user_id) VALUES($1, $2) RETURNING id;", eventID, userID).Scan(&eventID)
	if err != nil {
		return err
	}
	return nil
}

func (table *EventUserTable) DeleteConnection(userID, eventID int32) error {
	events, err := table.GetEvents(userID)
	if err != nil {
		return err
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	check := false
	for _, event := range events {
		if event.ID == eventID {
			check = true
		}
	}
	if !check {
		return ErrConnectionDoesntExist
	}
	_, err = table.Connection.Exec(context.Background(), "delete from \"EventUserModel\"  where event_id=$1 and user_id=$2", eventID, userID)
	if err != nil {
		return err
	}
	return nil
}

func NewEventTable(pool *pgxpool.Pool) *EventTable {
	return &EventTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}

func NewEventUserTable(pool *pgxpool.Pool) *EventUserTable {
	return &EventUserTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}
