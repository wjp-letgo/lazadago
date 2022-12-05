package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingListResult struct{
    Data	FreeShippingListDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]FreeShippingListDetailResponseEntity	`json:"detail"`
}
func (g FreeShippingListResult) String() string {
    return lib.ObjectToString(g)
}
type FreeShippingListDetailResponseEntity struct{}
