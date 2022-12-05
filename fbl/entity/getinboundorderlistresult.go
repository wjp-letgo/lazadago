package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInboundOrderListResult struct{
    Result	GetInboundOrderListResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetInboundOrderListDetailResponseEntity	`json:"detail"`
}
func (g GetInboundOrderListResult) String() string {
    return lib.ObjectToString(g)
}
type GetInboundOrderListDetailResponseEntity struct{}
