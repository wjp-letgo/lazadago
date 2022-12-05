package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWarehouseStockV3UnsellableResponseEntity struct {
	Available int `json:"available"`
	Reserved  int `json:"reserved"`
}

func (g GetWarehouseStockV3UnsellableResponseEntity) String() string {
	return lib.ObjectToString(g)
}
