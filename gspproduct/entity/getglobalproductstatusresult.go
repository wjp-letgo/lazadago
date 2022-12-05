package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetGlobalProductStatusResult struct{
    Data	string	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetGlobalProductStatusDetailResponseEntity	`json:"detail"`
}
func (g GetGlobalProductStatusResult) String() string {
    return lib.ObjectToString(g)
}
type GetGlobalProductStatusDetailResponseEntity struct{}
