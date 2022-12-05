package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SubmitSellerReplyResult struct{
    Data	bool	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SubmitSellerReplyDetailResponseEntity	`json:"detail"`
}
func (g SubmitSellerReplyResult) String() string {
    return lib.ObjectToString(g)
}
type SubmitSellerReplyDetailResponseEntity struct{}
