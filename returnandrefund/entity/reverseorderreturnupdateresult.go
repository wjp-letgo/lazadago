package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ReverseOrderReturnUpdateResult struct{
    Data	ReverseOrderReturnUpdateDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]ReverseOrderReturnUpdateDetailResponseEntity	`json:"detail"`
}
func (g ReverseOrderReturnUpdateResult) String() string {
    return lib.ObjectToString(g)
}
type ReverseOrderReturnUpdateDetailResponseEntity struct{}
