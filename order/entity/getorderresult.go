package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderResult struct{
    Data	GetOrderDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetOrderDetailResponseEntity	`json:"detail"`
}
func (g GetOrderResult) String() string {
    return lib.ObjectToString(g)
}

type GetOrderDetailResponseEntity struct{
    
}
