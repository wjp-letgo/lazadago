package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryLazadaBigbagInfoResult struct{
    Result	QueryLazadaBigbagInfoResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]QueryLazadaBigbagInfoDetailResponseEntity	`json:"detail"`
}
func (g QueryLazadaBigbagInfoResult) String() string {
    return lib.ObjectToString(g)
}
type QueryLazadaBigbagInfoDetailResponseEntity struct{}
