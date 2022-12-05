package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateSellerResult struct{
    Success	bool	`json:"success"`
    Data	string	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]UpdateSellerDetailResponseEntity	`json:"detail"`
}
func (g UpdateSellerResult) String() string {
    return lib.ObjectToString(g)
}
type UpdateSellerDetailResponseEntity struct{}
