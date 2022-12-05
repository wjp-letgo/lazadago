package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RemoveVideoResult struct{
    Success	bool	`json:"success"`
    ResultCode	string	`json:"result_code"`
    ResultMessage	string	`json:"result_message"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RemoveVideoDetailResponseEntity	`json:"detail"`
}
func (g RemoveVideoResult) String() string {
    return lib.ObjectToString(g)
}
type RemoveVideoDetailResponseEntity struct{}
