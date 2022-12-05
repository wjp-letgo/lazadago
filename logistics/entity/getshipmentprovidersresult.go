package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShipmentProvidersResult struct{
    Data	GetShipmentProvidersDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetShipmentProvidersDetailResponseEntity	`json:"detail"`
}
func (g GetShipmentProvidersResult) String() string {
    return lib.ObjectToString(g)
}
type GetShipmentProvidersDetailResponseEntity struct{}
