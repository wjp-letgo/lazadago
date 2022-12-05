package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadSellerBagNumberResult struct{
    Result	UploadSellerBagNumberResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UploadSellerBagNumberDetailResponseEntity	`json:"detail"`
}
func (g UploadSellerBagNumberResult) String() string {
    return lib.ObjectToString(g)
}
type UploadSellerBagNumberDetailResponseEntity struct{}
