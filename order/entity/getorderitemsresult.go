package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderItemsResult struct{
    Data	[]GetOrderItemsDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetOrderItemsDetailResponseEntity	`json:"detail"`
}
func (g GetOrderItemsResult) String() string {
    return lib.ObjectToString(g)
}

type GetOrderItemsDetailResponseEntity struct{
    
}
