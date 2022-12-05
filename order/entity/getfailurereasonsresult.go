package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFailureReasonsResult struct{
    Data	[]GetFailureReasonsDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetFailureReasonsDetailResponseEntity	`json:"detail"`
}
func (g GetFailureReasonsResult) String() string {
    return lib.ObjectToString(g)
}
type GetFailureReasonsDetailResponseEntity struct{
    
}