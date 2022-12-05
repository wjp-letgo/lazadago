package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToSOFFailedDeliveryResult struct{
    Data	SetStatusToSOFFailedDeliveryDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetStatusToSOFFailedDeliveryDetailResponseEntity	`json:"detail"`
}
func (g SetStatusToSOFFailedDeliveryResult) String() string {
    return lib.ObjectToString(g)
}

type SetStatusToSOFFailedDeliveryDetailResponseEntity struct{
    
}