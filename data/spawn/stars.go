package spawn

import (
	//"fmt"
	"math/rand"

	"github.com/sore0159/star_system/data"
)

func GenerateStarSystem(r *rand.Rand, step int) []*data.Star {
	stars := make(map[data.Location]*data.Star, 20)
	for i := 0; i < 20; i += 1 {
		sign := rand.Intn(8)
		s := data.NewStar(
			//fmt.Sprintf("Star %5d", r.Intn(99999)),
			r.Int63(), r.Int63(), r.Int63(),
		)
		switch sign {
		case 0:
		case 1:
			s.Y *= -1
		case 2:
			s.Z *= -1
		case 3:
			s.Y *= -1
			s.Z *= -1
		case 4:
			s.X *= -1
		case 5:
			s.X *= -1
			s.Y *= -1
		case 6:
			s.X *= -1
			s.Z *= -1
		case 7:
			s.Z *= -1
			s.X *= -1
			s.Y *= -1
		}
		stars[s.Location] = s
	}
	starList := make([]*data.Star, 0, len(stars))
	for _, s := range stars {
		starList = append(starList, s)
	}
	return starList
}
