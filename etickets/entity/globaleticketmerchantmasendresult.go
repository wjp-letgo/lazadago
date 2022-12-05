package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaSendResult struct{
    RespBody	GlobalEticketMerchantMaSendRespBodyResponseEntity	`json:"resp_body"`
    RetCode	string	`json:"ret_code"`
    RetMsg	string	`json:"ret_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GlobalEticketMerchantMaSendDetailResponseEntity	`json:"detail"`
}
func (g GlobalEticketMerchantMaSendResult) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaSendDetailResponseEntity struct{}
type GlobalEticketMerchantMaSendRespBodyResponseEntity struct{}