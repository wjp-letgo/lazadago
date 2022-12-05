package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateProductResult struct{
    Data	CreateProductDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]CreateProductDetailResponseEntity	`json:"detail"`
}
func (g CreateProductResult) String() string {
    return lib.ObjectToString(g)
}

type CreateProductDetailResponseEntity struct{}