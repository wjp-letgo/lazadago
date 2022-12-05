package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWarehouseStockStoreStocksResponseEntity struct {
	StoreCode string                                `json:"store_code"`
	Stocks    GetWarehouseStockStocksResponseEntity `json:"stocks"`
}

func (g GetWarehouseStockStoreStocksResponseEntity) String() string {
	return lib.ObjectToString(g)
}
