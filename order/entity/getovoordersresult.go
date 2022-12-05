package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOVOOrdersResult struct{
    Result	GetOVOOrdersResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetOVOOrdersDetailResponseEntity	`json:"detail"`
}
func (g GetOVOOrdersResult) String() string {
    return lib.ObjectToString(g)
}

type GetOVOOrdersDetailResponseEntity struct{
    
}
