package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockV3Result struct{
    Data	[]GetWarehouseStockV3DataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetWarehouseStockV3DetailResponseEntity	`json:"detail"`
}
func (g GetWarehouseStockV3Result) String() string {
    return lib.ObjectToString(g)
}
type GetWarehouseStockV3DetailResponseEntity struct{}
