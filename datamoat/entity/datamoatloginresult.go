package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DataMoatLoginResult struct{
    Result	DataMoatLoginResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]DataMoatLoginDetailResponseEntity	`json:"detail"`
}
func (g DataMoatLoginResult) String() string {
    return lib.ObjectToString(g)
}
type DataMoatLoginDetailResponseEntity struct{}
