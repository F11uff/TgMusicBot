package postgresql

import (
	"database/sql"
	"errors"
)

type Database struct {
	psql *sql.DB
}

func NewDatabase(psql *sql.DB) *Database {
	return &Database{psql: psql}
}

func (db *Database) Connect(urlConnect string) (*Database, error) {
	var err error

	db.psql, err = sql.Open("postgres", urlConnect)

	if err != nil {
		return nil, errors.New("fail to connection to database")
	}

	if err = db.psql.Ping(); err != nil {
		return nil, errors.New("fail to ping database")
	}

	return db, nil
}

func (db *Database) Insert() error {
	return nil
}
