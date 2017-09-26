package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/sore0159/star_system/data"
	"github.com/sore0159/star_system/data/spawn"
	"github.com/sore0159/star_system/pgdb"
)

func SpawnStars(db *pgdb.DB, steps int) error {
	const fileName = "STAR_DATA.txt"
	if steps > 0 {
		if err := CreateStarFile(fileName, steps); err != nil {
			return fmt.Errorf("file creation fail: %v", err)
		}
	}
	if err := ReadStarFile(db, fileName); err != nil {
		return fmt.Errorf("file load fail: %v", err)
	}
	return nil
}

func CreateStarFile(fileName string, steps int) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	write := func(stars []*data.Star) error {
		for i, s := range stars {
			if _, err = f.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\n", s.X, s.Y, s.Z, s.Name)); err != nil {
				return fmt.Errorf("file write error on star %d: %v", i, err)
			}
		}
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for step := 0; step < steps; step += 1 {
		stars := spawn.GenerateStarSystem(r, step)
		if err = write(stars); err != nil {
			return err
		}
	}
	return nil
}
func ReadStarFile(db *pgdb.DB, fileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getwd fail: %v", err)
	}
	fullN := filepath.Join(dir, fileName)

	if _, err := db.Exec("TRUNCATE stars CASCADE"); err != nil {
		return fmt.Errorf("truncation fail: %v", err)
	}
	if _, err := db.Exec(fmt.Sprintf("COPY stars FROM '%s'", fullN)); err != nil {
		return fmt.Errorf("copy from %s fail: %v", fullN, err)
	}
	return nil
}
