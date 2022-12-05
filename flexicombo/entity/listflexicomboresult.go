package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboResult struct{
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Data	ListFlexiComboDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]ListFlexiComboDetailResponseEntity	`json:"detail"`
}
func (g ListFlexiComboResult) String() string {
    return lib.ObjectToString(g)
}
type ListFlexiComboDetailResponseEntity struct{}
