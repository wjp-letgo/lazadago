package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderReasonListResult struct{
    Data	[]GetReverseOrderReasonListDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetReverseOrderReasonListDetailResponseEntity	`json:"detail"`
}
func (g GetReverseOrderReasonListResult) String() string {
    return lib.ObjectToString(g)
}
type GetReverseOrderReasonListDetailResponseEntity struct{}
