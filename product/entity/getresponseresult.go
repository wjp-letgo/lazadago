package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetResponseResult struct{
    Data	GetResponseDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetResponseDetailResponseEntity	`json:"detail"`
}
func (g GetResponseResult) String() string {
    return lib.ObjectToString(g)
}
type GetResponseDetailResponseEntity struct{}