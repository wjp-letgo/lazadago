package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetGlobalProductStatusParamsRequestEntity struct {
	SellerSku string `json:"sellerSku"`
}

func (g GetGlobalProductStatusParamsRequestEntity) String() string {
	return lib.ObjectToString(g)
}
