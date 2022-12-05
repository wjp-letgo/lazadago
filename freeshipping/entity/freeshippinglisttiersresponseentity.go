package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type FreeShippingListTiersResponseEntity struct {
	Filter string `json:"filter"`
	Result string `json:"result"`
}

func (g FreeShippingListTiersResponseEntity) String() string {
	return lib.ObjectToString(g)
}
