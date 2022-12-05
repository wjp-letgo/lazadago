package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateUserResult struct{
    Success	bool	`json:"success"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UpdateUserDetailResponseEntity	`json:"detail"`
}
func (g UpdateUserResult) String() string {
    return lib.ObjectToString(g)
}
type UpdateUserDetailResponseEntity struct{}
