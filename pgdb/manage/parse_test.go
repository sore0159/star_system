package main

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}
func TestSCHEMA(t *testing.T) {
	tables, err := ParseTableSQL(SCHEMA_FILENAME)
	log.Printf("SCHEMA TEST: %v, %v\n", tables, err)
}
