package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerMetricsByIdResult struct{
    Data	GetSellerMetricsByIdDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetSellerMetricsByIdDetailResponseEntity	`json:"detail"`
}
func (g GetSellerMetricsByIdResult) String() string {
    return lib.ObjectToString(g)
}
type GetSellerMetricsByIdDetailResponseEntity struct{}
