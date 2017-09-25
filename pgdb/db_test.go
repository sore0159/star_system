package pgdb

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestDB(t *testing.T) {
	db, err := LoadDB(DB_USER, DB_PASSWORD, DB_DB)
	log.Println("GOT: ", db == nil, err)
}
