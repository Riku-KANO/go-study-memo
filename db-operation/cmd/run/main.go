package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Riku-KANO/db-operation/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/lib/pq"
)

type App struct {
	server  Server
	errLog  *log.Logger
	infoLog *log.Logger
	models  models.Models
}

type Server struct {
	host string
	port string
	url  string
}


func (a *App) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", a.healthCheck)
	r.Get("/users", a.getUsersHandler)
	r.Post("/users", a.createUserHandler)
	return r
}

func (a *App) ListenAndServe() error {
	host := fmt.Sprintf("%s:%s", a.server.host, a.server.port)

	srv := http.Server{
		Handler: a.Routes(),
		Addr: host,
		ReadTimeout: 300 * time.Second,
	}

	a.infoLog.Printf("Server Listening on :%s\n", host)

	return srv.ListenAndServe()
}

func openDB(settings db.ConnectionURL) (db.Session, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		return nil, err
	}

	if err := sess.Ping(); err != nil {
		return nil, err
	}
	log.Println("DB connection is completed")
	return sess, nil
}

func main() {
	initdb := flag.Bool("initdb", false, "flag for initializing database. Default is false")
	flag.Parse()
	
	server := Server {
		host: "localhost",
		port: "8080",
		url: "http://localhost:8080",
	}

	var dbSettings = postgresql.ConnectionURL{
		Database: "database",
		Host:     "localhost:5455",
		User:     "root",
		Password: "password",
	}

	sess, err := openDB(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// initialization of database table
	if *initdb {
		query, err := os.ReadFile("./migrations/init.sql")
		if err != nil {
			log.Fatal(err)
		}

		_, err = sess.SQL().Exec((string(query)))
		if err != nil {
			log.Fatal(err)
		}

		return 
	}

	app := &App {
		server: server,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog: log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile),
		models: models.New(sess),
	}

	

	if err := app.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
