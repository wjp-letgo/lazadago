package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryAddressInformaitonResult struct{
    Result	QueryAddressInformaitonResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]QueryAddressInformaitonDetailResponseEntity	`json:"detail"`
}
func (g QueryAddressInformaitonResult) String() string {
    return lib.ObjectToString(g)
}
type QueryAddressInformaitonDetailResponseEntity struct{}
