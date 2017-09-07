package captain

import (
	"errors"
	"math/big"
)

type Captain struct {
	UID  big.Int `json:"uid"`
	Name string  `json:"name"`
}

type Academy interface {
	NewCaptain() (*Captain, error)
	SearchCaptain(uid *big.Int) (*Captain, error)
}

var ERR_CAP404 = errors.New("captain not found")
