package stars

import (
	"fmt"
	"math/rand"
)

func GenerateStarSystem(seed int64) []*Star {
	r := rand.New(rand.NewSource(seed))
	stars := make(map[Location]*Star, 20)
	for i := 0; i < 20; i += 1 {
		s := NewStar(
			fmt.Sprintf("Star %5d", r.Intn(99999)),
			r.Uint64(), r.Uint64(), r.Uint64(),
		)
		stars[s.Location] = s
	}
	starList := make([]*Star, 0, len(stars))
	for _, s := range stars {
		starList = append(starList, s)
	}
	return starList
}
