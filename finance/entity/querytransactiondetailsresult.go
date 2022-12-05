package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryTransactionDetailsResult struct{
    Data	[]QueryTransactionDetailsDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]QueryTransactionDetailsDetailResponseEntity	`json:"detail"`
}
func (g QueryTransactionDetailsResult) String() string {
    return lib.ObjectToString(g)
}
type QueryTransactionDetailsDetailResponseEntity struct{}
