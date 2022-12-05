package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaConsumeResult struct{
    RespBody	GlobalEticketMerchantMaConsumeRespBodyResponseEntity	`json:"resp_body"`
    RetCode	string	`json:"ret_code"`
    RetMsg	string	`json:"ret_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GlobalEticketMerchantMaConsumeDetailResponseEntity	`json:"detail"`
}
func (g GlobalEticketMerchantMaConsumeResult) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaConsumeDetailResponseEntity struct{}
