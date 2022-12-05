package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductReviewListResult struct{
    Data	GetProductReviewListDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetProductReviewListDetailResponseEntity	`json:"detail"`
}
func (g GetProductReviewListResult) String() string {
    return lib.ObjectToString(g)
}
type GetProductReviewListDetailResponseEntity struct{}
