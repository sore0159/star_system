package data

import (
	"math"
)

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

type StarPath [2]Location

func (sp StarPath) Len() float64 {
	return sp[0].Dist(sp[1])
}
func (sp StarPath) Same(sp2 StarPath) bool {
	return (sp[0] == sp2[0] && sp[1] == sp2[1]) ||
		(sp[0] == sp2[1] && sp[1] == sp2[0])
}
