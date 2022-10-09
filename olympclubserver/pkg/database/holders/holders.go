package holders

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type HolderTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type Holder struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func NewHolderTable(pool *pgxpool.Pool) *HolderTable {
	return &HolderTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}

func (table *HolderTable) InsertHolder(holder Holder) error {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	err := table.Connection.QueryRow(context.Background(), "insert into \"HolderModel\" (name, logo) VALUES($1, $2) RETURNING holder_id;", holder.Name, holder.Logo).Scan(&holder.ID)
	fmt.Println(err)
	return err
}

func (table *HolderTable) GetHolders(holderID int32) ([]Holder, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	holders := []Holder{}
	request := "Select * from \"HolderModel\" WHERE holder_id = $1"
	rows, err := table.Connection.Query(context.Background(), request, holderID)
	if err != nil {
		return []Holder{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		holder := Holder{}
		holder.ID = values[0].(int32)
		holder.Name = values[1].(string)
		holder.Logo = "img/" + values[2].(string)
		holders = append(holders, holder)
	}
	return holders, nil
}

func (table *HolderTable) GetAllHolders() ([]Holder, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	holders := []Holder{}
	request := "Select * from \"HolderModel\""
	rows, err := table.Connection.Query(context.Background(), request)
	if err != nil {
		return []Holder{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		holder := Holder{}
		holder.ID = values[0].(int32)
		holder.Name = values[1].(string)
		holder.Logo = "img/" + values[2].(string)
		holders = append(holders, holder)
	}
	return holders, nil
}
