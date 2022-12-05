package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockV3DataResponseEntity struct{
    FulfilmentSku	string	`json:"fulfilment_sku"`
    StoreStocks	[]GetWarehouseStockV3StoreStocksResponseEntity	`json:"store_stocks"`
}
func (g GetWarehouseStockV3DataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
