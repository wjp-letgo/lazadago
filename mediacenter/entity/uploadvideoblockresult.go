package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadVideoBlockResult struct{
    Success	bool	`json:"success"`
    ResultCode	string	`json:"result_code"`
    ETag	string	`json:"e_tag"`
    ResultMessage	string	`json:"result_message"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UploadVideoBlockDetailResponseEntity	`json:"detail"`
}
func (g UploadVideoBlockResult) String() string {
    return lib.ObjectToString(g)
}
type UploadVideoBlockDetailResponseEntity struct{}
