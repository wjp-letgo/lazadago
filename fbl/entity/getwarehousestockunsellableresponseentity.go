package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWarehouseStockUnsellableResponseEntity struct {
	Available int `json:"available"`
	Reserved  int `json:"reserved"`
}

func (g GetWarehouseStockUnsellableResponseEntity) String() string {
	return lib.ObjectToString(g)
}
