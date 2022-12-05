package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdatePriceQuantityResult struct{
    Data	UpdatePriceQuantityDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UpdatePriceQuantityDetailResponseEntity	`json:"detail"`
}
func (g UpdatePriceQuantityResult) String() string {
    return lib.ObjectToString(g)
}

type UpdatePriceQuantityDetailResponseEntity struct{}
type UpdatePriceQuantityDataResponseEntity struct{}