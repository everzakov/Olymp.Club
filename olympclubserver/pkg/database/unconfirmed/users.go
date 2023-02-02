package unconfirmed

import (
	user_model "OlympClub/pkg/database/users"
	"OlympClub/pkg/utils"
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UnConfirmedUsersTable struct {
	Connection *pgxpool.Pool
	mtx        *sync.Mutex
}

type UnConfirmedUser struct {
	ID        int32
	Email     string
	Token1    string
	Token2    string
	PassHash  string
	Confirmed bool
}

var (
	ErrUserDoesntExists       = fmt.Errorf("User is not exists")
	ErrUserExists             = fmt.Errorf("User exists")
	ErrUserIsAlreadyConfirmed = fmt.Errorf("User is already confirmed")
	ErrProblemWithTokens      = fmt.Errorf("Problem with tokens in Database")
)

// проверяем подтверждён ли пользователь
func (table *UnConfirmedUsersTable) IsUserConfirmed(email string) (bool, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UnConfirmedUsers\" where email=$1 and confirmed=$2;", email, true)
	if err != nil {
		return false, err
	}
	users := []UnConfirmedUser{}
	for rows.Next() {
		// переводим данные в список
		values, err := rows.Values()
		if err != nil {
			return false, err
		}
		unConfirmedUser := UnConfirmedUser{}
		unConfirmedUser.ID = values[0].(int32)
		unConfirmedUser.Email = values[1].(string)
		unConfirmedUser.Token1 = values[2].(string)
		unConfirmedUser.Token2 = values[3].(string)
		unConfirmedUser.PassHash = values[4].(string)
		unConfirmedUser.Confirmed = values[5].(bool)
		users = append(users, unConfirmedUser)
	}
	return len(users) > 0, nil
}

// создать неподтверждённого пользователя
func (table *UnConfirmedUsersTable) CreateUnconfirmedUser(email, pass_hash string) (UnConfirmedUser, error) {
	confirmed, err := table.IsUserConfirmed(email)
	table.mtx.Lock()
	defer table.mtx.Unlock()
	if err != nil {
		return UnConfirmedUser{}, err
	}
	if confirmed {
		return UnConfirmedUser{}, ErrUserExists
	}
	token1 := utils.GenerateSecureToken()
	token2 := utils.GenerateSecureToken()
	user := UnConfirmedUser{
		Email:     email,
		Token1:    token1,
		Token2:    token2,
		PassHash:  pass_hash,
		Confirmed: false,
	}
	err = table.Connection.QueryRow(context.Background(), "insert into \"UnConfirmedUsers\" (email, token1, token2, pass_hash, confirmed) VALUES($1, $2, $3, $4, $5) RETURNING id;", user.Email, user.Token1, user.Token2, user.PassHash, false).Scan(&user.ID)
	if err != nil {
		return UnConfirmedUser{}, err
	}
	return user, nil
}

// получить пользователь по токенам
func (table *UnConfirmedUsersTable) GetUsersByTokens(token1, token2 string) ([]UnConfirmedUser, error) {
	table.mtx.Lock()
	defer table.mtx.Unlock()
	rows, err := table.Connection.Query(context.Background(), "select * from \"UnConfirmedUsers\" where token1=$1 and token2=$2;", token1, token2)
	if err != nil {
		return []UnConfirmedUser{}, err
	}
	users := []UnConfirmedUser{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return []UnConfirmedUser{}, err
		}
		unConfirmedUser := UnConfirmedUser{}
		// convert DB types to Go types
		unConfirmedUser.ID = values[0].(int32)
		unConfirmedUser.Email = values[1].(string)
		unConfirmedUser.Token1 = values[2].(string)
		unConfirmedUser.Token2 = values[3].(string)
		unConfirmedUser.PassHash = values[4].(string)
		unConfirmedUser.Confirmed = values[5].(bool)
		users = append(users, unConfirmedUser)
	}
	if len(users) == 0 {
		return []UnConfirmedUser{}, ErrUserDoesntExists
	}
	return users, nil
}

// подтвердить пользователя
func (table *UnConfirmedUsersTable) ConfirmUser(token1, token2 string) (user_model.User, error) {
	users, err := table.GetUsersByTokens(token1, token2)
	table.mtx.Lock()
	defer table.mtx.Unlock()
	if err != nil {
		return user_model.User{}, err
	}
	if len(users) == 0 {
		return user_model.User{}, ErrUserDoesntExists
	}
	if users[0].Confirmed {
		return user_model.User{}, ErrUserIsAlreadyConfirmed
	}
	unConfirmedUser := users[0]
	user := user_model.User{
		Email:    unConfirmedUser.Email,
		Token1:   unConfirmedUser.Token1,
		Token2:   unConfirmedUser.Token2,
		PassHash: unConfirmedUser.PassHash,
	}
	rows, err := table.Connection.Query(context.Background(), "update \"UnConfirmedUsers\" set confirmed=$1 where id=$2", true, unConfirmedUser.ID)
	if err != nil {
		return user_model.User{}, err
	}
	rows.Close()
	err = table.Connection.QueryRow(context.Background(), "insert into \"UserModel\" (email, token1, token2, pass_hash) VALUES($1, $2, $3, $4) RETURNING id;", user.Email, user.Token1, user.Token2, user.PassHash).Scan(&user.ID)
	if err != nil {
		return user_model.User{}, err
	}
	return user, nil
}

// удалить пользователя
func (table *UnConfirmedUsersTable) DeleteUser(token1, token2 string) error {
	users, err := table.GetUsersByTokens(token1, token2)
	table.mtx.Lock()
	defer table.mtx.Unlock()
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return ErrUserDoesntExists
	}
	if users[0].Confirmed {
		return ErrUserIsAlreadyConfirmed
	}
	_, err = table.Connection.Exec(context.Background(), "delete from \"UnConfirmedUsers\" where token1=$1 and token2=$2 and confirmed=$3", token1, token2, false)
	return err
}

func NewUnConfirmedUsersTable(pool *pgxpool.Pool) *UnConfirmedUsersTable {
	return &UnConfirmedUsersTable{
		Connection: pool,
		mtx:        &sync.Mutex{},
	}
}
