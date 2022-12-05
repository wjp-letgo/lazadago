package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerItemLimitResult struct{
    Success	bool	`json:"success"`
    ErrorCodes	[]string	`json:"errorCodes"`
    ErrorMsgs	[]string	`json:"errorMsgs"`
    Data	GetSellerItemLimitDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetSellerItemLimitDetailResponseEntity	`json:"detail"`
}
func (g GetSellerItemLimitResult) String() string {
    return lib.ObjectToString(g)
}

type GetSellerItemLimitDetailResponseEntity struct{}