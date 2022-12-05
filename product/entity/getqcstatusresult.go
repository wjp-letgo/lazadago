package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetQcStatusResult struct{
    Data	[]GetQcStatusDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetQcStatusDetailResponseEntity	`json:"detail"`
}
func (g GetQcStatusResult) String() string {
    return lib.ObjectToString(g)
}
type GetQcStatusDetailResponseEntity struct{}