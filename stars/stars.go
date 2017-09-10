package stars

import (
	"errors"
	"math"
)

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

type Location struct {
	X, Y, Z uint64
}

func Diff(a, b uint64) uint64 {
	if a < b {
		return b - a
	}
	return a - b
}

func (l Location) Dist(l2 Location) float64 {
	dx, dy, dz := float64(Diff(l.X, l2.X)), float64(Diff(l.Y, l2.Y)), float64(Diff(l.Z, l2.Z))
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (s Star) Dist(s2 Star) float64 {
	return s.Location.Dist(s2.Location)
}
func (s Star) PathTo(s2 Star) StarPath {
	return StarPath([2]Location{s.Location, s2.Location})
}

type StarPath [2]Location

func (sp StarPath) Len() float64 {
	return sp[0].Dist(sp[1])
}
func (sp StarPath) Same(sp2 StarPath) bool {
	return (sp[0] == sp2[0] && sp[1] == sp2[1]) ||
		(sp[0] == sp2[1] && sp[1] == sp2[0])
}

type StarSystem interface {
	CreateStars([]*Star) error
	SearchStar(l Location) (*Star, error)
	LocalStars(Location) ([]*Star, error)

	ValidPath(path StarPath) (bool, error)
	CheckBlazed(path StarPath) (bool, error)
	SetBlazed(path StarPath) error
}

var ERR_STAR404 = errors.New("star not found")
