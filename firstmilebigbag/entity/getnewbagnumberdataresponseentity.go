package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetNewBagNumberDataResponseEntity struct {
	LzdBagNumber string `json:"lzd_bag_number"`
}

func (g GetNewBagNumberDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
