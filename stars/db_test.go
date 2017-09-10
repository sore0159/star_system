package stars

import (
	"log"
	"testing"
)

func TestDB(t *testing.T) {
	log.Println("TEST DB")
	m := new(MockDB)
	stars := []*Star{NewStar("star1", 0, 0, 0), NewStar("star2", 1, 1, 1)}
	m.CreateStars(stars)
	log.Printf("Made stars: %#v\n", stars)

	m.Save("TEST_SAVE.json")
	m2, err := LoadMockDB("TEST_SAVE.json")
	if err != nil {
		log.Fatalf("Err loading mockdb: %v\n", err)
	}
	s1, err := m2.SearchStar(stars[0].Location)
	log.Printf("Search results: %#v, %v", s1, err)
	s2, err := m2.SearchStar(Location{0, 1, 0})
	log.Printf("Search results: %#v, %v", s2, err)

}
