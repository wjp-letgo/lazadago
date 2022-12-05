package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteFlexiComboProductsResult struct{
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]DeleteFlexiComboProductsDetailResponseEntity	`json:"detail"`
}
func (g DeleteFlexiComboProductsResult) String() string {
    return lib.ObjectToString(g)
}
type DeleteFlexiComboProductsDetailResponseEntity struct{}
