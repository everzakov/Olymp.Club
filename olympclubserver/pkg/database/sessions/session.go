package sessions

import (
	"OlympClub/pkg/utils"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type SessionModel struct {
	Connection *pgxpool.Pool
}

type Session struct {
	ID     int32
	UserID int32
	Token  string
	Time   time.Time
}

func (table *SessionModel) GetSessions(token string) ([]Session, error) {
	sessions := []Session{}
	queryString := "Select * from \"SessionModel\" where token=$1"
	rows, err := table.Connection.Query(context.Background(), queryString, token)
	if err != nil {
		return []Session{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		session := Session{}
		session.ID = values[0].(int32)
		session.UserID = values[1].(int32)
		session.Token = values[2].(string)
		session.Time = values[3].(time.Time)
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (table *SessionModel) GetSessionsByUserID(userID int32) ([]Session, error) {
	sessions := []Session{}
	queryString := "Select * from \"SessionModel\" where user_id=$1"
	rows, err := table.Connection.Query(context.Background(), queryString, userID)
	if err != nil {
		return []Session{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		session := Session{}
		session.ID = values[0].(int32)
		session.UserID = values[1].(int32)
		session.Token = values[2].(string)
	}
	return sessions, nil
}

func (table *SessionModel) CreateSession(userID int32) (Session, error) {
	token := utils.GenerateSecureToken()
	session := Session{
		UserID: userID,
		Token:  token,
	}
	err := table.Connection.QueryRow(context.Background(), "insert into \"SessionModel\" (user_id, Token) VALUES($1, $2) RETURNING id;", session.UserID, session.Token).Scan(&session.ID)
	if err != nil {
		return Session{}, err
	}
	return session, nil
}

func NewSessionModel(pool *pgxpool.Pool) *SessionModel {
	return &SessionModel{
		Connection: pool,
	}
}
