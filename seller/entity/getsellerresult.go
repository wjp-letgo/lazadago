package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerResult struct{
    Data	GetSellerDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetSellerDetailResponseEntity	`json:"detail"`
}
func (g GetSellerResult) String() string {
    return lib.ObjectToString(g)
}
type GetSellerDetailResponseEntity struct{}
