package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type RetailFulfilmentCreateDataResponseEntity struct {
	ScItemId      int64  `json:"sc_item_id"`
	FulfilmentSku string `json:"fulfilment_sku"`
}

func (g RetailFulfilmentCreateDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
