package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type FreeShippingUpdateTiersRequestEntity struct {
	Filter string `json:"filter"`
	Result string `json:"result"`
}

func (g FreeShippingUpdateTiersRequestEntity) String() string {
	return lib.ObjectToString(g)
}
