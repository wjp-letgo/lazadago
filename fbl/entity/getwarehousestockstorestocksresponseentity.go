package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockStoreStocksResponseEntity struct{
    StoreCode	string	`json:"store_code"`
    Stocks	GetWarehouseStockStocksResponseEntity	`json:"stocks"`
}
func (g GetWarehouseStockStoreStocksResponseEntity) String() string {
    return lib.ObjectToString(g)
}
