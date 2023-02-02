package admins

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AdminTable struct {
	Connection *pgxpool.Pool
}

type Admin struct {
	ID       int32
	UserID   int32
	Priority int32
}

// получить список админов
func (table *AdminTable) GetAdmins() ([]Admin, error) {
	admins := []Admin{}
	rows, err := table.Connection.Query(context.Background(), "Select * from \"AdminModel\"")
	if err != nil {
		return []Admin{}, err
	}
	for rows.Next() {
		// переводим данные в список админов
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		admin := Admin{}
		admin.ID = values[0].(int32)
		admin.UserID = values[1].(int32)
		admin.Priority = values[2].(int32)
		admins = append(admins, admin)
	}
	return admins, nil
}

// проверка что пользователь - админ
func (table *AdminTable) CheckAdmin(userID int32) (bool, error) {
	admins := []Admin{}
	rows, err := table.Connection.Query(context.Background(), "Select * from \"AdminModel\" where user_id=$1", userID)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		admin := Admin{}
		admin.ID = values[0].(int32)
		admin.UserID = values[1].(int32)
		admin.Priority = values[2].(int32)
		admins = append(admins, admin)
	}
	return len(admins) > 0, nil
}

func NewAdminTable(dbpool *pgxpool.Pool) *AdminTable {
	return &AdminTable{
		Connection: dbpool,
	}
}
