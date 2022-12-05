package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategoryAttributesResult struct{
    Data	[]GetCategoryAttributesDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetCategoryAttributesDetailResponseEntity	`json:"detail"`
}
func (g GetCategoryAttributesResult) String() string {
    return lib.ObjectToString(g)
}
type GetCategoryAttributesDetailResponseEntity struct{}