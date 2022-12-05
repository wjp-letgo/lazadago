package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockV3StoreStocksResponseEntity struct{
    StoreCode	string	`json:"store_code"`
    Stocks	GetWarehouseStockV3StocksResponseEntity	`json:"stocks"`
}
func (g GetWarehouseStockV3StoreStocksResponseEntity) String() string {
    return lib.ObjectToString(g)
}
