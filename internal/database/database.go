package database

import (
	"database/sql"
	"fmt"
)

var (
	conn *connectionOptions
)

// connectionOptions describes necessary information
// to access a Postgres instance
type connectionOptions struct {
	host     string
	port     string
	database string
	username string
	password string
}

// String converts connection options to a format
// that the database library understands
func (c *connectionOptions) string() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.host, c.port, c.username, c.password, c.database,
	)
}

// Init stores the database connection options
func Init(host, port, database, username, password string) {
	conn = &connectionOptions{
		host, port, database, username, password,
	}
}

// Connect tries to connect to the database
// using options that describe the necessary
// information
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", conn.string())
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Ping tries to connect to the database
// and pings this to test if the connection is okay
func Ping() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
