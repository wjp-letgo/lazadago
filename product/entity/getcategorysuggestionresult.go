package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategorySuggestionResult struct{
    Data	GetCategorySuggestionDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetCategorySuggestionDetailResponseEntity	`json:"detail"`
}
func (g GetCategorySuggestionResult) String() string {
    return lib.ObjectToString(g)
}
type GetCategorySuggestionDetailResponseEntity struct{}