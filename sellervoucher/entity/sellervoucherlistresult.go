package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherListResult struct{
    Data	SellerVoucherListDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SellerVoucherListDetailResponseEntity	`json:"detail"`
}
func (g SellerVoucherListResult) String() string {
    return lib.ObjectToString(g)
}
type SellerVoucherListDetailResponseEntity struct{}
