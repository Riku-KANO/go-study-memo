package db_operation

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Riku-KANO/db-operation/models"
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

	r.Get("/users", a.getUsersHandler)

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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	initdb := flag.Bool("initdb", false, "flag for initializing database. Default is false")
	flag.Parse()
	
	server := Server {
		host: "localhost",
		port: "8080",
		url: "http://localhost:8080",
	}

	const (
		user = "app"
		password = "password"
		host = "localhost"
		port = "5455"
		dbname = "database"
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialization of database table
	if *initdb {
		query, err := os.ReadFile("./migrations/init.sql")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(string(query))
		if err != nil {
			log.Fatal(err)
		}

		return 
	}

	app := &App {
		server: server,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog: log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile),
		models: models.New(),
	}

	

	if err := app.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
