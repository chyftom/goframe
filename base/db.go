package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "192.168.9.128"
	port     = 5432
	user     = "postgres"
	password = "postgrespassword"
	dbname   = "postgres"
)


var db *sql.DB
var psqlInfo string

func init() {
	var err error = nil

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

}

func DB () *sql.DB {
	return db
}

func NewDB() *sql.DB {
	var err error = nil
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Failed to new db:", err)
	}

	return db
}