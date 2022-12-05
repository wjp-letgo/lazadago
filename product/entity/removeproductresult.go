package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RemoveProductResult struct{
    Data	RemoveProductDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RemoveProductDetailResponseEntity	`json:"detail"`
}
func (g RemoveProductResult) String() string {
    return lib.ObjectToString(g)
}
type RemoveProductDetailResponseEntity struct{}
type RemoveProductDataResponseEntity struct{}