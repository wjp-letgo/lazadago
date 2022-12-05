package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingGetResult struct{
    Data	FreeShippingGetDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]FreeShippingGetDetailResponseEntity	`json:"detail"`
}
func (g FreeShippingGetResult) String() string {
    return lib.ObjectToString(g)
}
type FreeShippingGetDetailResponseEntity struct{}
