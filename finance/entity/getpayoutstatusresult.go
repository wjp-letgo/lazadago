package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutStatusResult struct{
    Data	[]GetPayoutStatusDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetPayoutStatusDetailResponseEntity	`json:"detail"`
}
func (g GetPayoutStatusResult) String() string {
    return lib.ObjectToString(g)
}
type GetPayoutStatusDetailResponseEntity struct{}
