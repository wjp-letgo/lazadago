package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockV3StocksResponseEntity struct{
    Sellable	GetWarehouseStockV3SellableResponseEntity	`json:"sellable"`
    Unsellable	GetWarehouseStockV3UnsellableResponseEntity	`json:"unsellable"`
    Pending	GetWarehouseStockV3PendingResponseEntity	`json:"pending"`
    Transfer	GetWarehouseStockV3TransferResponseEntity	`json:"transfer"`
}
func (g GetWarehouseStockV3StocksResponseEntity) String() string {
    return lib.ObjectToString(g)
}
