package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateFlexiComboResult struct{
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UpdateFlexiComboDetailResponseEntity	`json:"detail"`
}
func (g UpdateFlexiComboResult) String() string {
    return lib.ObjectToString(g)
}
type UpdateFlexiComboDetailResponseEntity struct{}
