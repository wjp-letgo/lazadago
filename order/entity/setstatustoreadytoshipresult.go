package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToReadyToShipResult struct{
    Data	SetStatusToReadyToShipDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetStatusToReadyToShipDetailResponseEntity	`json:"detail"`
}
func (g SetStatusToReadyToShipResult) String() string {
    return lib.ObjectToString(g)
}

type SetStatusToReadyToShipDetailResponseEntity struct{
    
}