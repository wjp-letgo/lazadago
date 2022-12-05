package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeactivateFlexiComboResult struct{
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]DeactivateFlexiComboDetailResponseEntity	`json:"detail"`
}
func (g DeactivateFlexiComboResult) String() string {
    return lib.ObjectToString(g)
}
type DeactivateFlexiComboDetailResponseEntity struct{}
