package users

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserModel struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

var (
	ErrUserDoesntExists = fmt.Errorf("User is not exists")
)

type User struct {
	ID       int32
	Email    string
	PassHash string
	Token1   string
	Token2   string
}

func (table *UserModel) IsUserExists(email string) (bool, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UserModel\" where email=$1;", email)
	if err != nil {
		return false, err
	}
	users := []User{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return false, err
		}
		user := User{}
		// convert DB types to Go types
		user.ID = values[0].(int32)
		user.Email = values[1].(string)
		user.PassHash = values[2].(string)
		user.Token1 = values[3].(string)
		user.Token2 = values[4].(string)
		users = append(users, user)
	}
	return len(users) > 0, nil
}

func (table *UserModel) GetUsersByEmailAndPassword(email, pass_hash string) ([]User, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UserModel\" where email=$1 and pass_hash=$2;", email, pass_hash)
	if err != nil {
		return []User{}, err
	}
	users := []User{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return []User{}, err
		}
		user := User{}
		// convert DB types to Go types
		user.ID = values[0].(int32)
		user.Email = values[1].(string)
		user.PassHash = values[2].(string)
		user.Token1 = values[3].(string)
		user.Token2 = values[4].(string)
		users = append(users, user)
	}
	if len(users) == 0 {
		return []User{}, ErrUserDoesntExists
	}
	return users, nil
}

func (table *UserModel) GetUsersByEmail(email string) ([]User, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UserModel\" where email=$1;", email)
	if err != nil {
		return []User{}, err
	}
	users := []User{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return []User{}, err
		}
		user := User{}
		// convert DB types to Go types
		user.ID = values[0].(int32)
		user.Email = values[1].(string)
		user.PassHash = values[2].(string)
		user.Token1 = values[3].(string)
		user.Token2 = values[4].(string)
		users = append(users, user)
	}
	if len(users) == 0 {
		return []User{}, ErrUserDoesntExists
	}
	return users, nil
}

func (table *UserModel) CheckUserByEmailAndPassword(email, pass_hash string) (bool, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UserModel\" where email=$1 and pass_hash=$2;", email, pass_hash)
	if err != nil {
		return false, err
	}
	users := []User{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return false, err
		}
		user := User{}
		// convert DB types to Go types
		user.ID = values[0].(int32)
		user.Email = values[1].(string)
		user.PassHash = values[2].(string)
		user.Token1 = values[3].(string)
		user.Token2 = values[4].(string)
		users = append(users, user)
	}
	if len(users) == 0 {
		return false, ErrUserDoesntExists
	}
	return len(users) > 0, nil
}

func (table *UserModel) GetUsersByTokens(token1, token2 string) ([]User, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UserModel\" where token1=$1 and token2=$2;", token1, token2)
	if err != nil {
		return []User{}, err
	}
	users := []User{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return []User{}, err
		}
		user := User{}
		// convert DB types to Go types
		user.ID = values[0].(int32)
		user.Email = values[1].(string)
		user.PassHash = values[2].(string)
		user.Token1 = values[3].(string)
		user.Token2 = values[4].(string)
		users = append(users, user)
	}
	if len(users) == 0 {
		return []User{}, ErrUserDoesntExists
	}
	return users, nil
}

func (table *UserModel) UpdatePassword(token1, token2, pass_hash string) error {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	_, err := table.Connection.Query(context.Background(), "update \"UserModel\" set pass_hash=$1 where token1=$2 and token2=$3", pass_hash, token1, token2)
	return err
}

func NewUserModel(pool *pgxpool.Pool) *UserModel {
	return &UserModel{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}
