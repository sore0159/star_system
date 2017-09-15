package data

import (
	"math/big"
)

type Captain struct {
	UID  big.Int `json:"uid"`
	Name string  `json:"name"`
}
