package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWarehouseStockSellableResponseEntity struct {
	Available int `json:"available"`
	Reserved  int `json:"reserved"`
}

func (g GetWarehouseStockSellableResponseEntity) String() string {
	return lib.ObjectToString(g)
}
