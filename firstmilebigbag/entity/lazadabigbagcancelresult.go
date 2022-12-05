package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCancelResult struct{
    Result	LazadaBigbagCancelResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]LazadaBigbagCancelDetailResponseEntity	`json:"detail"`
}
func (g LazadaBigbagCancelResult) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagCancelDetailResponseEntity struct{}
