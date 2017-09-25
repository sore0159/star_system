package spawn

import (
	"fmt"
	"math/rand"

	"github.com/sore0159/star_system/data"
)

func GenerateStarSystem(r *rand.Rand, step int) []*data.Star {
	stars := make(map[data.Location]*data.Star, 20)
	for i := 0; i < 20; i += 1 {
		s := data.NewStar(
			fmt.Sprintf("Star %5d", r.Intn(99999)),
			r.Uint64(), r.Uint64(), r.Uint64(),
		)
		stars[s.Location] = s
	}
	starList := make([]*data.Star, 0, len(stars))
	for _, s := range stars {
		starList = append(starList, s)
	}
	return starList
}
