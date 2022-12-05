package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CompleteCreateVideoResult struct{
    Success	bool	`json:"success"`
    ResultCode	string	`json:"result_code"`
    VideoId	string	`json:"video_id"`
    ResultMessage	string	`json:"result_message"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]CompleteCreateVideoDetailResponseEntity	`json:"detail"`
}
func (g CompleteCreateVideoResult) String() string {
    return lib.ObjectToString(g)
}
type CompleteCreateVideoDetailResponseEntity struct{}
