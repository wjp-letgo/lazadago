package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RedeemOrderItemsResult struct{
    Data	RedeemOrderItemsDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RedeemOrderItemsDetailResponseEntity	`json:"detail"`
}
func (g RedeemOrderItemsResult) String() string {
    return lib.ObjectToString(g)
}
type RedeemOrderItemsDetailResponseEntity struct{}
