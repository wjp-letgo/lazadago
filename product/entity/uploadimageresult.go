package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadImageResult struct{
    Data	UploadImageDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UploadImageDetailResponseEntity	`json:"detail"`
}
func (g UploadImageResult) String() string {
    return lib.ObjectToString(g)
}
type UploadImageDetailResponseEntity struct{}