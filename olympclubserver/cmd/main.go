package main

import (
	"OlympClub/pkg/database/admins"
	"OlympClub/pkg/database/events"
	"OlympClub/pkg/database/holders"
	"OlympClub/pkg/database/news"
	"OlympClub/pkg/database/olympiads"
	"OlympClub/pkg/database/sessions"
	"OlympClub/pkg/database/unconfirmed"
	"OlympClub/pkg/database/users"
	admin_handler "OlympClub/pkg/handlers/admins"
	event_handler "OlympClub/pkg/handlers/events"
	"OlympClub/pkg/types"

	auth_handler "OlympClub/pkg/handlers/auth"
	holder_handler "OlympClub/pkg/handlers/holders"
	olympiad_handler "OlympClub/pkg/handlers/olympiads"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	s := r.PathPrefix("/api/v1/").Subrouter()

	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	mailInfo := types.MailInfo{
		From:     os.Getenv("MAIL_FROM"),
		User:     os.Getenv("MAIL_USER"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Addr:     os.Getenv("MAIL_ADDR"),
		Host:     os.Getenv("MAIL_HOST"),
		FrontURL: os.Getenv("FRONT_URL"),
	}

	olympiadModel := olympiads.NewOlympiadTable(dbpool)
	bigOlympiadModel := olympiads.NewBigOlympiadTable(dbpool)
	olympiadUserModel := olympiads.NewOlympiadUserTable(dbpool)

	userModel := users.NewUserModel(dbpool)
	unConfirmedUsers := unconfirmed.NewUnConfirmedUsersTable(dbpool)

	holdersModel := holders.NewHolderTable(dbpool)
	sessionModel := sessions.NewSessionModel(dbpool)

	eventModel := events.NewEventTable(dbpool)
	eventUserModel := events.NewEventUserTable(dbpool)

	adminModel := admins.NewAdminTable(dbpool)

	newsModel := news.NewNewsTable(dbpool)

	authHandler := auth_handler.AuthHandler{
		UnConfirmedUsersTable: unConfirmedUsers,
		UsersModel:            userModel,
		SessionModel:          sessionModel,
		MailInfo:              mailInfo,
	}
	olympiad_handler := olympiad_handler.OlympiadHandler{
		OlympiadModel:     olympiadModel,
		BigOlympiadModel:  bigOlympiadModel,
		SessionModel:      sessionModel,
		OlympiadUserModel: olympiadUserModel,
		NewsModel:         newsModel,
	}
	holder_handler := holder_handler.HolderHandler{
		HolderTable: holdersModel,
	}
	eventHandler := event_handler.EventHandler{
		EventModel:     eventModel,
		EventUserModel: eventUserModel,
		SessionModel:   sessionModel,
		NewsModel:      newsModel,
	}

	adminHandler := admin_handler.AdminHandler{
		AdminModel:       adminModel,
		HolderTable:      holdersModel,
		SessionModel:     sessionModel,
		BigOlympiadModel: bigOlympiadModel,
		OlympiadModel:    olympiadModel,
		EventModel:       eventModel,
		NewsModel:        newsModel,
	}

	authHandler.RegisterHandler(s)
	olympiad_handler.RegisterHandler(s)
	holder_handler.RegisterHandler(s)
	eventHandler.RegisterHandler(s)
	adminHandler.RegisterHandler(s)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
