package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var count int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Couldn't Connect to Postgres")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {

		return nil, err

	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	connection, err := openDB(dsn)
	for {
		if err != nil {
			log.Println("Postgres not yet ready...")
			count++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if count > 10 {
			log.Panicln(err)
		}

		log.Println("Backing off for 2 secs...")
		time.Sleep(2 * time.Second)
		continue
	}
}
