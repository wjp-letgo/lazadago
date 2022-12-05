package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RetailFulfilmentUpdateResult struct{
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Success	bool	`json:"success"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RetailFulfilmentUpdateDetailResponseEntity	`json:"detail"`
}
func (g RetailFulfilmentUpdateResult) String() string {
    return lib.ObjectToString(g)
}
type RetailFulfilmentUpdateDetailResponseEntity struct{}