package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFlexiComboDetailsResult struct{
    Data	GetFlexiComboDetailsDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetFlexiComboDetailsDetailResponseEntity	`json:"detail"`
}
func (g GetFlexiComboDetailsResult) String() string {
    return lib.ObjectToString(g)
}
type GetFlexiComboDetailsDetailResponseEntity struct{}
