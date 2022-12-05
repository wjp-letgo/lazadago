package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerPerformanceResult struct{
    Data	GetSellerPerformanceDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetSellerPerformanceDetailResponseEntity	`json:"detail"`
}
func (g GetSellerPerformanceResult) String() string {
    return lib.ObjectToString(g)
}
type GetSellerPerformanceDetailResponseEntity struct{}
