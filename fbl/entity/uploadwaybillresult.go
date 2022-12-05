package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadWaybillResult struct{
    Success	bool	`json:"success"`
    ErrorMessage	string	`json:"error_message"`
    ErrorCode	string	`json:"error_code"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UploadWaybillDetailResponseEntity	`json:"detail"`
}
func (g UploadWaybillResult) String() string {
    return lib.ObjectToString(g)
}
type UploadWaybillDetailResponseEntity struct{}
