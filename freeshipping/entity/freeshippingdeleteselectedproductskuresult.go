package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingDeleteSelectedProductSKUResult struct{
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]FreeShippingDeleteSelectedProductSKUDetailResponseEntity	`json:"detail"`
}
func (g FreeShippingDeleteSelectedProductSKUResult) String() string {
    return lib.ObjectToString(g)
}
type FreeShippingDeleteSelectedProductSKUDetailResponseEntity struct{}
