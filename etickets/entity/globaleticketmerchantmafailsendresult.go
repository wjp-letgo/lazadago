package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaFailsendResult struct{
    RespBody	GlobalEticketMerchantMaFailsendRespBodyResponseEntity	`json:"resp_body"`
    RetCode	string	`json:"ret_code"`
    RetMsg	string	`json:"ret_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GlobalEticketMerchantMaFailsendDetailResponseEntity	`json:"detail"`
}
func (g GlobalEticketMerchantMaFailsendResult) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaFailsendDetailResponseEntity struct{}

type GlobalEticketMerchantMaFailsendRespBodyResponseEntity struct{}