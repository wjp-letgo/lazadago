package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetReverseOrderDetailProductDTOResponseEntity struct {
	ProductId int    `json:"product_id"`
	Sku       string `json:"sku"`
}

func (g GetReverseOrderDetailProductDTOResponseEntity) String() string {
	return lib.ObjectToString(g)
}
