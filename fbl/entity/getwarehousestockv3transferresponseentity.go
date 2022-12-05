package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWarehouseStockV3TransferResponseEntity struct {
	Available int `json:"available"`
	Reserved  int `json:"reserved"`
}

func (g GetWarehouseStockV3TransferResponseEntity) String() string {
	return lib.ObjectToString(g)
}
