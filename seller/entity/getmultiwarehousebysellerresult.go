package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMultiWarehouseBySellerResult struct{
    Result	GetMultiWarehouseBySellerResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetMultiWarehouseBySellerDetailResponseEntity	`json:"detail"`
}
func (g GetMultiWarehouseBySellerResult) String() string {
    return lib.ObjectToString(g)
}
type GetMultiWarehouseBySellerDetailResponseEntity struct{}
