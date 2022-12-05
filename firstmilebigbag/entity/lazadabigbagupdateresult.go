package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagUpdateResult struct{
    Result	LazadaBigbagUpdateResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]LazadaBigbagUpdateDetailResponseEntity	`json:"detail"`
}
func (g LazadaBigbagUpdateResult) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagUpdateDetailResponseEntity struct{}
