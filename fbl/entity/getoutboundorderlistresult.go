package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOutboundOrderListResult struct{
    Result	GetOutboundOrderListResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetOutboundOrderListDetailResponseEntity	`json:"detail"`
}
func (g GetOutboundOrderListResult) String() string {
    return lib.ObjectToString(g)
}
type GetOutboundOrderListDetailResponseEntity struct{}
