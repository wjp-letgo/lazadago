package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrdersForSellerResult struct{
    Result	GetReverseOrdersForSellerResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetReverseOrdersForSellerDetailResponseEntity	`json:"detail"`
}
func (g GetReverseOrdersForSellerResult) String() string {
    return lib.ObjectToString(g)
}
type GetReverseOrdersForSellerDetailResponseEntity struct{}
