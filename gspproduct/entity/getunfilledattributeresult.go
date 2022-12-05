package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeResult struct{
    Data	GetUnfilledAttributeDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorDetail	string	`json:"error_detail"`
    ErrorCode	string	`json:"error_code"`
    Errors	string	`json:"errors"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetUnfilledAttributeDetailResponseEntity	`json:"detail"`
}
func (g GetUnfilledAttributeResult) String() string {
    return lib.ObjectToString(g)
}
type GetUnfilledAttributeDetailResponseEntity struct{}
