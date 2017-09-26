package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/sore0159/star_system/data"
	"github.com/sore0159/star_system/pgdb"
)

func TestScanDB(t *testing.T) {
	log.Println("Testing DB Rows")
	db, err := pgdb.LoadDB(pgdb.DB_USER, pgdb.DB_PASSWORD, pgdb.DB_DB)
	if err != nil {
		log.Printf("DB error: %v\n", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM stars LIMIT 10")
	if err != nil {
		log.Printf("QUERY error: %v\n", err)
		return
	}
	var count int
	defer rows.Close()
	for rows.Next() {
		count += 1
		var s data.Star
		err = rows.Scan(&s.X, &s.Y, &s.Z, &s.Name)
		if err != nil {
			log.Printf("ROWS SCAN ERR: %v\n", err)
		} else {
			fmt.Printf("Scanned: %v\n", s)
		}
	}
	if err = rows.Err(); err != nil {
		log.Printf("ROWS CLOSE ERR: %v\n", err)
	} else {
		log.Printf("Scanned %d rows\n", count)
	}
}
