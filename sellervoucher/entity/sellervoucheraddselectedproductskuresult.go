package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherAddSelectedProductSKUResult struct{
    Data	SellerVoucherAddSelectedProductSKUDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SellerVoucherAddSelectedProductSKUDetailResponseEntity	`json:"detail"`
}
func (g SellerVoucherAddSelectedProductSKUResult) String() string {
    return lib.ObjectToString(g)
}
type SellerVoucherAddSelectedProductSKUDetailResponseEntity struct{}

type SellerVoucherAddSelectedProductSKUDataResponseEntity struct{
    SkuId string `json:"sku id"`
}
func (g SellerVoucherAddSelectedProductSKUDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}