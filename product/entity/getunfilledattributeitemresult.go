package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeItemResult struct{
    Success	bool	`json:"success"`
    TotalProducts	int	`json:"total_products"`
    Products	[]GetUnfilledAttributeItemProductsResponseEntity	`json:"products"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetUnfilledAttributeItemDetailResponseEntity	`json:"detail"`
}
func (g GetUnfilledAttributeItemResult) String() string {
    return lib.ObjectToString(g)
}
type GetUnfilledAttributeItemDetailResponseEntity struct{}