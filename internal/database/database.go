package database

import (
	"database/sql"
	"fmt"
)

// DB is the database reference that other
// packages can use to interact with the
// database
var DB *sql.DB

// ConnectionOptions describes necessary information
// to access a Postgres instance
type ConnectionOptions struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

// String converts connection options to a format
// that the database library understands
func (c *ConnectionOptions) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.Database,
	)
}

// Connect tries to connect to the database
// using options that describe the necessary
// information
func Connect(connOptions *ConnectionOptions) error {
	var err error

	DB, err = sql.Open("postgres", connOptions.String())
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	return nil
}
