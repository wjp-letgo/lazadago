package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateGlobalProductResult struct{
    Data	CreateGlobalProductDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]CreateGlobalProductDetailResponseEntity	`json:"detail"`
}
func (g CreateGlobalProductResult) String() string {
    return lib.ObjectToString(g)
}
type CreateGlobalProductDetailResponseEntity struct{}
