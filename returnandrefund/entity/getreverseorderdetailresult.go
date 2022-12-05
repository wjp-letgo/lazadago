package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderDetailResult struct{
    Data	GetReverseOrderDetailDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetReverseOrderDetailDetailResponseEntity	`json:"detail"`
}
func (g GetReverseOrderDetailResult) String() string {
    return lib.ObjectToString(g)
}
type GetReverseOrderDetailDetailResponseEntity struct{}
