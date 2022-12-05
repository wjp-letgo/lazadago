package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaAvailableResult struct{
    RespBody	GlobalEticketMerchantMaAvailableRespBodyResponseEntity	`json:"resp_body"`
    RetCode	string	`json:"ret_code"`
    RetMsg	string	`json:"ret_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GlobalEticketMerchantMaAvailableDetailResponseEntity	`json:"detail"`
}
func (g GlobalEticketMerchantMaAvailableResult) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaAvailableDetailResponseEntity struct{}
