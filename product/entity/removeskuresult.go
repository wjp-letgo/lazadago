package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RemoveSkuResult struct{
    Data	RemoveSkuDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RemoveSkuDetailResponseEntity	`json:"detail"`
}
func (g RemoveSkuResult) String() string {
    return lib.ObjectToString(g)
}
type RemoveSkuDetailResponseEntity struct{}
type RemoveSkuDataResponseEntity struct{}