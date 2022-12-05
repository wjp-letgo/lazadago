package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInboundOrderDetailResult struct{
    Data	GetInboundOrderDetailDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetInboundOrderDetailDetailResponseEntity	`json:"detail"`
}
func (g GetInboundOrderDetailResult) String() string {
    return lib.ObjectToString(g)
}
type GetInboundOrderDetailDetailResponseEntity struct{}
