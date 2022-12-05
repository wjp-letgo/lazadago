package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type InitReverseOrderCancelResult struct{
    Data	InitReverseOrderCancelDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]InitReverseOrderCancelDetailResponseEntity	`json:"detail"`
}
func (g InitReverseOrderCancelResult) String() string {
    return lib.ObjectToString(g)
}
type InitReverseOrderCancelDetailResponseEntity struct{}
