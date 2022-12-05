package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFulfillmentProductDetailResult struct{
    Data	[]GetFulfillmentProductDetailDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetFulfillmentProductDetailDetailResponseEntity	`json:"detail"`
}
func (g GetFulfillmentProductDetailResult) String() string {
    return lib.ObjectToString(g)
}
type GetFulfillmentProductDetailDetailResponseEntity struct{}
