package pgdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func LoadDB(user, pass, dbName string) (mdb *DB, err error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbName))
	if err != nil {
		return nil, fmt.Errorf("loaddb failure: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("loaddb ping failure: %v", err)
	}
	return &DB{db}, nil
}
