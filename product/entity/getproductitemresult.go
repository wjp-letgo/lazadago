package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductItemResult struct{
    Data	GetProductItemDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetProductItemDetailResponseEntity	`json:"detail"`
}
func (g GetProductItemResult) String() string {
    return lib.ObjectToString(g)
}
type GetProductItemDetailResponseEntity struct{}