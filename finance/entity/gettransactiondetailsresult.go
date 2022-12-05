package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTransactionDetailsResult struct{
    Data	[]GetTransactionDetailsDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetTransactionDetailsDetailResponseEntity	`json:"detail"`
}
func (g GetTransactionDetailsResult) String() string {
    return lib.ObjectToString(g)
}
type GetTransactionDetailsDetailResponseEntity struct{}
