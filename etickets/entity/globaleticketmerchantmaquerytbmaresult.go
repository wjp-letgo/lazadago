package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaQueryTbMaResult struct{
    RespBody	GlobalEticketMerchantMaQueryTbMaRespBodyResponseEntity	`json:"resp_body"`
    RetCode	string	`json:"ret_code"`
    RetMsg	string	`json:"ret_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GlobalEticketMerchantMaQueryTbMaDetailResponseEntity	`json:"detail"`
}
func (g GlobalEticketMerchantMaQueryTbMaResult) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaQueryTbMaDetailResponseEntity struct{}
type GlobalEticketMerchantMaQueryTbMaRespBodyResponseEntity struct{}