package data

type Star struct {
	Location `json:"location"`
	Name     string `json:"name"`
}

func NewStar(name string, x, y, z uint64) *Star {
	return &Star{
		Name: name,
		Location: Location{
			X: x,
			Y: y,
			Z: z,
		},
	}
}

func (s Star) Dist(s2 Star) float64 {
	return s.Location.Dist(s2.Location)
}
func (s Star) PathTo(s2 Star) StarPath {
	return StarPath([2]Location{s.Location, s2.Location})
}
