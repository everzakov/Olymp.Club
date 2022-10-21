package olympiads

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BigOlympiadTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type BigOlympiad struct {
	ID          int32      `json:"id"`
	Name        string     `json:"name"`
	Short       string     `json:"short"`
	Logo        string     `json:"logo"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Olympiads   []Olympiad `json:"olympiads"`
}

func NewBigOlympiadTable(pool *pgxpool.Pool) *BigOlympiadTable {
	return &BigOlympiadTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}

var (
	ErrBigOlympiadIsAlreadyExisted = fmt.Errorf("Olympiad is already existed")
)

func (table *BigOlympiadTable) CreateBigOlympiad(olympiad BigOlympiad) (BigOlympiad, error) {
	filter := NewBigOlympiadFilter()
	filter.Short = olympiad.Short
	olympiads, err := table.GetBigOlympiads(filter)
	if err != nil {
		return BigOlympiad{}, err
	}
	if len(olympiads) != 0 {
		return BigOlympiad{}, ErrBigOlympiadIsAlreadyExisted
	}
	table.mtx.Lock()
	defer table.mtx.Unlock()
	_, err = table.Connection.Exec(context.Background(), "insert into \"BigOlympiadModel\" (name, short, logo, description, status) VALUES($1, $2, $3, $4, $5)", olympiad.Name, olympiad.Short, olympiad.Logo, olympiad.Description, olympiad.Status)
	// fmt.Println(err)
	if err != nil {
		return BigOlympiad{}, err
	}
	return olympiad, nil
}

func (table *BigOlympiadTable) GetBigOlympiads(filter BigOlympiadFilter) ([]BigOlympiad, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	olympiads := []BigOlympiad{}
	queryString, queryArgs := GetQueryBigOlympiadOptions(filter)
	rows, err := table.Connection.Query(context.Background(), queryString, queryArgs...)
	if err != nil {
		return olympiads, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		olympiad := BigOlympiad{}
		olympiad.ID = values[0].(int32)
		olympiad.Name = values[1].(string)
		olympiad.Short = values[2].(string)
		olympiad.Logo = "img/" + values[3].(string)
		olympiad.Description = values[4].(string)
		olympiad.Status = values[5].(string)
		olympiads = append(olympiads, olympiad)
	}
	return olympiads, nil
}
