package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaQueryResult struct{
    RespBody	GlobalEticketMerchantMaQueryRespBodyResponseEntity	`json:"resp_body"`
    RetCode	string	`json:"ret_code"`
    RetMsg	string	`json:"ret_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GlobalEticketMerchantMaQueryDetailResponseEntity	`json:"detail"`
}
func (g GlobalEticketMerchantMaQueryResult) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaQueryDetailResponseEntity struct{}
