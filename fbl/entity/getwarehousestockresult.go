package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockResult struct{
    Data	[]GetWarehouseStockDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetWarehouseStockDetailResponseEntity	`json:"detail"`
}
func (g GetWarehouseStockResult) String() string {
    return lib.ObjectToString(g)
}
type GetWarehouseStockDetailResponseEntity struct{}
