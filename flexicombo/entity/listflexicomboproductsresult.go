package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboProductsResult struct{
    Data	ListFlexiComboProductsDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]ListFlexiComboProductsDetailResponseEntity	`json:"detail"`
}
func (g ListFlexiComboProductsResult) String() string {
    return lib.ObjectToString(g)
}
type ListFlexiComboProductsDetailResponseEntity struct{}
