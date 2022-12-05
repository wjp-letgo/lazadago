package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddFlexiComboProductsResult struct{
    Data	AddFlexiComboProductsDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]AddFlexiComboProductsDetailResponseEntity	`json:"detail"`
}
func (g AddFlexiComboProductsResult) String() string {
    return lib.ObjectToString(g)
}
type AddFlexiComboProductsDetailResponseEntity struct{}
type AddFlexiComboProductsDataResponseEntity struct{
    SkuId string `json:"sku id"`
}
func (g AddFlexiComboProductsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}