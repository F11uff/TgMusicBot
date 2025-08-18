package storage

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

var psql2 *sql.DB

type Database struct {
	psql  *sql.DB
	Music *Music
	User  *User
}

func NewDatabase() *Database {
	return &Database{
		psql:  psql2,
		Music: NewMusic(),
		User:  NewUser(),
	}
}

func (db *Database) Connect(urlConnect string) (*Database, error) {
	var err error

	db.psql, err = sql.Open("postgres", urlConnect)

	if err != nil {
		return nil, errors.New("fail to connection to database")
	}

	if err = db.GetDB().Ping(); err != nil {
		return nil, errors.New("fail to ping database")
	}

	return db, nil
}

func (db *Database) GetDB() *sql.DB {
	return db.psql
}

func (db *Database) Close() error {
	if db.psql != nil {
		return db.psql.Close()
	}

	return nil
}

func (db *Database) Insert() error {
	return nil
}
