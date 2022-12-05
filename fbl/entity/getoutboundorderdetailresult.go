package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOutboundOrderDetailResult struct{
    Data	GetOutboundOrderDetailDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetOutboundOrderDetailDetailResponseEntity	`json:"detail"`
}
func (g GetOutboundOrderDetailResult) String() string {
    return lib.ObjectToString(g)
}
type GetOutboundOrderDetailDetailResponseEntity struct{}
