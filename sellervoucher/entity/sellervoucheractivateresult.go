package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherActivateResult struct{
    Success	bool	`json:"success"`
    ErrorCode	int	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SellerVoucherActivateDetailResponseEntity	`json:"detail"`
}
func (g SellerVoucherActivateResult) String() string {
    return lib.ObjectToString(g)
}
type SellerVoucherActivateDetailResponseEntity struct{}
