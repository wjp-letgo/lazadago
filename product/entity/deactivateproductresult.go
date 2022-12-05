package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeactivateProductResult struct{
    Data	DeactivateProductDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]DeactivateProductDetailResponseEntity	`json:"detail"`
}
func (g DeactivateProductResult) String() string {
    return lib.ObjectToString(g)
}
type DeactivateProductDetailResponseEntity struct{}
type DeactivateProductDataResponseEntity struct{}