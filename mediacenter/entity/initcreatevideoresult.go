package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type InitCreateVideoResult struct{
    UploadId	string	`json:"upload_id"`
    Success	bool	`json:"success"`
    ResultCode	string	`json:"result_code"`
    ResultMessage	string	`json:"result_message"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]InitCreateVideoDetailResponseEntity	`json:"detail"`
}
func (g InitCreateVideoResult) String() string {
    return lib.ObjectToString(g)
}
type InitCreateVideoDetailResponseEntity struct{}
