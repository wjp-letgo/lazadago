package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCommitResult struct{
    Result	LazadaBigbagCommitResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]LazadaBigbagCommitDetailResponseEntity	`json:"detail"`
}
func (g LazadaBigbagCommitResult) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagCommitDetailResponseEntity struct{}
