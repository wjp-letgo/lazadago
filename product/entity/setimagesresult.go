package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetImagesResult struct{
    Data	SetImagesDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetImagesDetailResponseEntity	`json:"detail"`
}
func (g SetImagesResult) String() string {
    return lib.ObjectToString(g)
}

type SetImagesDataResponseEntity struct{}
type SetImagesDetailResponseEntity struct{}