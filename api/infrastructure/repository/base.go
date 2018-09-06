package repository

import "database/sql"

type (
	// Execer is interface of sql.DB.Exec methods.
	Execer interface {
		Exec(string, ...interface{}) (sql.Result, error)
	}
	// Queryer is interface of sql.DB.Query and sql.DB.QueryRow methods.
	Queryer interface {
		Query(string, ...interface{}) (*sql.Rows, error)
		QueryRow(string, ...interface{}) *sql.Row
	}
	// Preparer is interface of sql.DB.Prepare methods.
	Preparer interface {
		Prepare(string) (*sql.Stmt, error)
	}
	// Beginner is interface of sql.DB.Begin methods.
	Beginner interface {
		Begin() (*sql.Tx, error)
	}
)

// DB is interface of *sql.DB
type DB interface {
	Execer
	Queryer
	Preparer
}

// DBConnector is a interface od *sql.DB
type DBConnector interface {
	DB
	Beginner
}

// Tx is interface of *sql.Tx
type Tx interface {
	Rollback() error
	Commit() error
	DB
}
