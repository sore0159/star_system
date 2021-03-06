package main

import (
	"log"
	"os"

	"github.com/sore0159/star_system/pgdb"
)

func main() {
	opts, err := ParseOpts()
	if err != nil {
		log.Printf("Parse Opts error: %v\n", err)
		if opts != nil && opts.HelpFlag {
			PrintHelp()
		}
		return
	}
	if opts.HelpFlag {
		PrintHelp()
		return
	}
	db, err := pgdb.LoadDB(NORMAL_DB_USER, NORMAL_DB_PASSWORD, NORMAL_DB_DB)
	if err != nil {
		log.Printf("DB error: %v\n", err)
		return
	}
	defer db.Close()
	if opts.DropFlag || opts.CreateFlag {
		if err := ManageTables(os.Stdout, db, opts); err != nil {
			log.Printf("MANAGE ERROR: %v\n", err)
			return
		}
	}
	if opts.SpawnFlag {
		// SpawnStars needs superdb to use COPY FROM file
		sdb, err := pgdb.LoadDB(SUPER_DB_USER, SUPER_DB_PASSWORD, SUPER_DB_DB)
		if err != nil {
			log.Printf("SUPER DB error: %v\n", err)
			return
		}
		defer sdb.Close()
		if opts.StarSteps > 0 {
			log.Println("Regenerating star system!")
		} else {
			log.Println("Reloading star system from file!")
		}
		if err = SpawnStars(sdb, opts.StarSteps); err != nil {
			log.Printf("Spawn err: %v\n", err)
			return
		}
	}
}
