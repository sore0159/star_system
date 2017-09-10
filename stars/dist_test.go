package stars

import (
	"log"
	"math"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	log.Println("TEST TWO")
	l1 := Location{X: math.MaxUint64, Y: math.MaxUint64, Z: math.MaxUint64}
	l2 := Location{}
	log.Printf("Dist: %f\n", l1.Dist(l2))
}
