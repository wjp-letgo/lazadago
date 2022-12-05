package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockStocksResponseEntity struct{
    Sellable	GetWarehouseStockSellableResponseEntity	`json:"sellable"`
    Unsellable	GetWarehouseStockUnsellableResponseEntity	`json:"unsellable"`
    Pending	GetWarehouseStockPendingResponseEntity	`json:"pending"`
}
func (g GetWarehouseStockStocksResponseEntity) String() string {
    return lib.ObjectToString(g)
}
