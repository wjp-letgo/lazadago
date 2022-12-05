package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingAddSelectedProductSKUResult struct{
    Data	FreeShippingAddSelectedProductSKUDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]FreeShippingAddSelectedProductSKUDetailResponseEntity	`json:"detail"`
}
func (g FreeShippingAddSelectedProductSKUResult) String() string {
    return lib.ObjectToString(g)
}
type FreeShippingAddSelectedProductSKUDetailResponseEntity struct{}

type FreeShippingAddSelectedProductSKUDataResponseEntity struct{
    SkuId string `json:"sku id"`
}
func (g FreeShippingAddSelectedProductSKUDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}