package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/cauabernardino/bucket-list/config"
	_ "github.com/lib/pq"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = errors.New("no matching record")

// Database represents a Database connection
type Database struct {
	Conn *sql.DB
}

// Connect opens a connection with Database
func Connect(username, password, database string) (Database, error) {
	db := Database{}

	conn, err := sql.Open("postgres", config.DBConnectString)
	if err != nil {
		return db, err
	}

	db.Conn = conn

	if err = db.Conn.Ping(); err != nil {
		return Database{}, err
	}

	log.Println("Database connection started")
	return db, nil
}
