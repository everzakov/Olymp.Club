package news

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type NewsTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type News struct {
	ID          int32
	Title       string
	Description string
	Table       string
	Key         int32
}

func (table *NewsTable) InsertNews(news News) error {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	// добавить новость
	err := table.Connection.QueryRow(context.Background(), "insert into \"NewsModel\" (\"Title\", \"Description\", \"TableStruct\", \"Key\") VALUES($1, $2, $3, $4) RETURNING id;", news.Title, news.Description, news.Table, news.Key).Scan(&news.ID)
	return err
}

func (table *NewsTable) GetNews(filter NewsFilter) ([]News, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()

	// переводим фильтр в query
	queryString, queryArgs := GetQueryNewsOptions(filter)
	rows, err := table.Connection.Query(context.Background(), queryString, queryArgs...)
	if err != nil {
		return []News{}, err
	}

	news := []News{}
	for rows.Next() {
		// перевод данных в список новостей
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		newsStruct := News{}
		newsStruct.ID = values[0].(int32)
		newsStruct.Title = values[1].(string)
		newsStruct.Description = values[2].(string)
		newsStruct.Table = values[3].(string)
		newsStruct.Key = values[4].(int32)
		news = append(news, newsStruct)
	}
	return news, nil
}

func NewNewsTable(dbpool *pgxpool.Pool) *NewsTable {
	return &NewsTable{
		Connection: dbpool,
		mtx:        &sync.Mutex{},
	}
}
