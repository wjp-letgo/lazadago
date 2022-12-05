package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RetailFulfilmentCreateResult struct{
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Data	RetailFulfilmentCreateDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RetailFulfilmentCreateDetailResponseEntity	`json:"detail"`
}
func (g RetailFulfilmentCreateResult) String() string {
    return lib.ObjectToString(g)
}
type RetailFulfilmentCreateDetailResponseEntity struct{}