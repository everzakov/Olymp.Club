package olympiads

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type OlympiadTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type OlympiadUserTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type Olympiad struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	Subject       string `json:"subject"`
	Level         string `json:"level"`
	Img           string `json:"img"`
	Short         string `json:"short"`
	BigOlympiadID int32  `json:"big_olympiad_id"`
	Status        string `json:"status"`
	Grade         string `json:"grade"`
	HolderID      int32  `json:"holder_id"`
	Website       string `json:"website"`
}

var (
	ErrOlympiadIsAlreadyExisted = fmt.Errorf("Olympiad is already existed")
	ErrConnectionExists         = fmt.Errorf("Connection already exists")
	ErrConnectionDoesntExist    = fmt.Errorf("Connection does not exist")
)

func (table *OlympiadTable) CreateOlympiad(olympiad Olympiad) (Olympiad, error) {
	filter := NewOlympiadFilter()
	filter.OlympiadShort = olympiad.Short
	filter.BigOlympiadID = olympiad.BigOlympiadID
	olympiads, err := table.GetOlympiads(filter)
	if err != nil {
		return Olympiad{}, err
	}
	if len(olympiads) != 0 {
		return Olympiad{}, ErrOlympiadIsAlreadyExisted
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	err = table.Connection.QueryRow(context.Background(), "insert into \"OlympiadModel\" (Name, Subject, Level, Img, Short, Big_Olympiad_ID, Status, Grade, Holder_Id, Website) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;", olympiad.Name, olympiad.Subject, olympiad.Level, olympiad.Img, olympiad.Short, olympiad.BigOlympiadID, olympiad.Status, olympiad.Grade, olympiad.HolderID, olympiad.Website).Scan(&olympiad.ID)
	// fmt.Println("lel", err)
	if err != nil {
		return Olympiad{}, err
	}
	return olympiad, nil
}

func (table *OlympiadTable) GetOlympiads(filter OlympiadFilter) ([]Olympiad, error) {
	olympiads := []Olympiad{}
	queryString, queryArgs := GetQueryOlympiadOptions(filter)
	rows, err := table.Connection.Query(context.Background(), queryString, queryArgs...)
	if err != nil {
		return []Olympiad{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		olympiad := Olympiad{}
		olympiad.ID = values[0].(int32)
		olympiad.Name = values[1].(string)
		olympiad.Subject = values[2].(string)
		olympiad.Level = values[3].(string)
		olympiad.Img = "img/" + values[4].(string)
		olympiad.Short = values[5].(string)
		olympiad.BigOlympiadID = values[6].(int32)
		olympiad.Status = values[7].(string)
		olympiad.Grade = values[8].(string)
		olympiad.HolderID = values[9].(int32)
		olympiad.Website = values[10].(string)
		olympiads = append(olympiads, olympiad)
	}
	return olympiads, nil
}

func (table *OlympiadUserTable) GetOlympiads(userID int32) ([]Olympiad, error) {
	rows, err := table.Connection.Query(context.Background(), "select om.* from \"OlympiadUserModel\" as ou left outer join \"OlympiadModel\" as om on ou.olympiad_id = om.id where ou.user_id=$1", userID)
	if err != nil {
		return []Olympiad{}, err
	}
	olympiads := []Olympiad{}
	for rows.Next() {
		values, err := rows.Values()
		// fmt.Println(err)
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		olympiad := Olympiad{}
		olympiad.ID = values[0].(int32)
		olympiad.Name = values[1].(string)
		olympiad.Subject = values[2].(string)
		olympiad.Level = values[3].(string)
		olympiad.Img = "img/" + values[4].(string)
		olympiad.Short = values[5].(string)
		olympiad.BigOlympiadID = values[6].(int32)
		olympiad.Status = values[7].(string)
		olympiad.Grade = values[8].(string)
		olympiad.HolderID = values[9].(int32)
		olympiad.Website = values[10].(string)
		olympiads = append(olympiads, olympiad)
	}
	return olympiads, nil
}

func (table *OlympiadUserTable) CreateConnection(userID, olympiadID int32) error {
	olympiads, err := table.GetOlympiads(userID)
	if err != nil {
		return err
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	check := false
	for _, olympiad := range olympiads {
		if olympiad.ID == olympiadID {
			check = true
		}
	}
	if check {
		return ErrConnectionExists
	}
	err = table.Connection.QueryRow(context.Background(), "insert into \"OlympiadUserModel\" (olympiad_id, user_id) VALUES($1, $2) RETURNING id;", olympiadID, userID).Scan(&olympiadID)
	if err != nil {
		return err
	}
	return nil
}

func (table *OlympiadUserTable) DeleteConnection(userID, olympiadID int32) error {
	olympiads, err := table.GetOlympiads(userID)
	if err != nil {
		return err
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	check := false
	for _, olympiad := range olympiads {
		if olympiad.ID == olympiadID {
			check = true
		}
	}
	if !check {
		return ErrConnectionDoesntExist
	}
	_, err = table.Connection.Exec(context.Background(), "delete from \"OlympiadUserModel\"  where event_id=$1 and user_id=$2", olympiadID, userID)
	if err != nil {
		return err
	}
	return nil
}

func NewOlympiadTable(pool *pgxpool.Pool) *OlympiadTable {
	return &OlympiadTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}

func NewOlympiadUserTable(pool *pgxpool.Pool) *OlympiadUserTable {
	return &OlympiadUserTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}
