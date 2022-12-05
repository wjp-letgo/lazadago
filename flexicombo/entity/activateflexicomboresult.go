package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ActivateFlexiComboResult struct{
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]ActivateFlexiComboDetailResponseEntity	`json:"detail"`
}
func (g ActivateFlexiComboResult) String() string {
    return lib.ObjectToString(g)
}
type ActivateFlexiComboDetailResponseEntity struct{}
