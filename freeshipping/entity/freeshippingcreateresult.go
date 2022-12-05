package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingCreateResult struct{
    Data	int	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]FreeShippingCreateDetailResponseEntity	`json:"detail"`
}
func (g FreeShippingCreateResult) String() string {
    return lib.ObjectToString(g)
}
type FreeShippingCreateDetailResponseEntity struct{}
