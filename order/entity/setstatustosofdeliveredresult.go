package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToSOFDeliveredResult struct{
    Data	SetStatusToSOFDeliveredDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetStatusToSOFDeliveredDetailResponseEntity	`json:"detail"`
}
func (g SetStatusToSOFDeliveredResult) String() string {
    return lib.ObjectToString(g)
}

type SetStatusToSOFDeliveredDetailResponseEntity struct{
    
}