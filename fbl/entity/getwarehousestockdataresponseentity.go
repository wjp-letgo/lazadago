package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWarehouseStockDataResponseEntity struct {
	FulfilmentSku string                                       `json:"fulfilment_sku"`
	StoreStocks   []GetWarehouseStockStoreStocksResponseEntity `json:"store_stocks"`
}

func (g GetWarehouseStockDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
