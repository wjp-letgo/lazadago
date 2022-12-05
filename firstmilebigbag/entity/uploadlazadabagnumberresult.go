package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadLazadaBagNumberResult struct{
    Result	UploadLazadaBagNumberResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UploadLazadaBagNumberDetailResponseEntity	`json:"detail"`
}
func (g UploadLazadaBagNumberResult) String() string {
    return lib.ObjectToString(g)
}
type UploadLazadaBagNumberDetailResponseEntity struct{}
