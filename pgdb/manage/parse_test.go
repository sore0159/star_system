package main

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}
func TestSchema(t *testing.T) {
	tables, err := ParseTableSQL(SCHEMA_FILENAME)
	log.Printf("SCHEMA TEST: %v, %v\n", tables, err)
}
