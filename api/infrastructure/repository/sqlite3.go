package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewSqlite3 create new sqlite3 connection.
func NewSqlite3(ds string) (DBConnector, error) {
	db, err := sql.Open("sqlite3", ds)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
