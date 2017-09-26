package data

import (
	"errors"
)

type Provider interface {
	Academy
	StarSystem
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

type Academy interface {
	NewCaptain() (*Captain, error)
	SearchCaptain(uid UID) (*Captain, error)
}

var ERR_CAP404 = errors.New("captain not found")
