package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingRegionsQueryResult struct{
    Data	[]FreeShippingRegionsQueryDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]FreeShippingRegionsQueryDetailResponseEntity	`json:"detail"`
}
func (g FreeShippingRegionsQueryResult) String() string {
    return lib.ObjectToString(g)
}
type FreeShippingRegionsQueryDetailResponseEntity struct{}
