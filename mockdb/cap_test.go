package mockdb

import (
	"log"
	"testing"

	"github.com/sore0159/star_system/data"
)

func TestONE(t *testing.T) {
	log.Println("TEST ONE")
}

func TestCaps(t *testing.T) {
	log.Println("TEST CAPS")
	m := NewMockProvider()
	c1, _ := m.NewCaptain()
	c2, _ := m.NewCaptain()
	log.Printf("Made captains: %#v, %#v\n", c1, c2)
	log.Printf("UID CMP: %v\n", c1.UID.Cmp(&c2.UID))
	m.Save("TEST_SAVE.json")
	m2, err := LoadMockProvider("TEST_SAVE.json")
	if err != nil {
		log.Fatalf("Err loading mockdb: %v\n", err)
	}
	c3, err := m2.SearchCaptain(&c1.UID)
	log.Printf("Search results: %#v, %v", c3, err)

}

func TestStars(t *testing.T) {
	log.Println("TEST STARS")
	m := NewMockProvider()
	stars := []*data.Star{data.NewStar("star1", 0, 0, 0), data.NewStar("star2", 1, 1, 1)}
	m.CreateStars(stars)
	log.Printf("Made stars: %#v\n", stars)

	m.Save("TEST_SAVE.json")
	m2, err := LoadMockProvider("TEST_SAVE.json")
	if err != nil {
		log.Fatalf("Err loading mockdb: %v\n", err)
	}
	s1, err := m2.SearchStar(stars[0].Location)
	log.Printf("Search results: %#v, %v", s1, err)
	s2, err := m2.SearchStar(data.Location{0, 1, 0})
	log.Printf("Search results: %#v, %v", s2, err)

}
