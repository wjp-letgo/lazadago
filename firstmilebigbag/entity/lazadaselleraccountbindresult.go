package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaSellerAccountBindResult struct{
    Result	LazadaSellerAccountBindResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]LazadaSellerAccountBindDetailResponseEntity	`json:"detail"`
}
func (g LazadaSellerAccountBindResult) String() string {
    return lib.ObjectToString(g)
}
type LazadaSellerAccountBindDetailResponseEntity struct{}
