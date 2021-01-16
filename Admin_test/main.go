package main

import (
	"awesomeProject3/handlers"
	"awesomeProject3/repository"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

func MustDBConn(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	if db.Ping() != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	conn := MustDBConn(dsn)

	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	router := mux.NewRouter()

	repo := repository.Repository{Conn: conn}

	router.Handle("/api/user/", handlers.GUserHandler{Conn: conn, Logger: logger, Repo: repo}).Methods(http.MethodGet)
	router.Handle("/api/login/", handlers.UserHandler{Conn: conn, Logger: logger, Repo: repo}).Methods(http.MethodPost)
	router.Handle("/api/register/", handlers.AUserHandler{Conn: conn, Logger: logger, Repo: repo}).Methods(http.MethodPost)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8009",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
