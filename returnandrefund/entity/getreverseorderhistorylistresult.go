package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderHistoryListResult struct{
    Data	GetReverseOrderHistoryListDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetReverseOrderHistoryListDetailResponseEntity	`json:"detail"`
}
func (g GetReverseOrderHistoryListResult) String() string {
    return lib.ObjectToString(g)
}
type GetReverseOrderHistoryListDetailResponseEntity struct{}
