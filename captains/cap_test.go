package captain

import (
	"log"
	"testing"
)

func TestONE(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	log.Println("TEST TWO")
	m := new(MockDB)
	c1, _ := m.NewCaptain()
	c2, _ := m.NewCaptain()
	log.Printf("Made captains: %#v, %#v\n", c1, c2)
	log.Printf("UID CMP: %v\n", c1.UID.Cmp(&c2.UID))
	m.Save("TEST_SAVE.json")
	m2, err := LoadMockDB("TEST_SAVE.json")
	if err != nil {
		log.Fatalf("Err loading mockdb: %v\n", err)
	}
	c3, err := m2.SearchCaptain(&c1.UID)
	log.Printf("Search results: %#v, %v", c3, err)

}
