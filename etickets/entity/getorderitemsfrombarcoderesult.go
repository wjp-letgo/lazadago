package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderItemsFromBarCodeResult struct{
    Data	GetOrderItemsFromBarCodeDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetOrderItemsFromBarCodeDetailResponseEntity	`json:"detail"`
}
func (g GetOrderItemsFromBarCodeResult) String() string {
    return lib.ObjectToString(g)
}
type GetOrderItemsFromBarCodeDetailResponseEntity struct{}
