package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategoryTreeResult struct{
    Data	[]GetCategoryTreeDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetCategoryTreeDetailResponseEntity	`json:"detail"`
}
func (g GetCategoryTreeResult) String() string {
    return lib.ObjectToString(g)
}
type GetCategoryTreeDetailResponseEntity struct{}

//GetCategoryTreeDataResponseEntity
type GetCategoryTreeDataResponseEntity struct{
    CategoryID int64 `json:"category_id"`
    Var bool `json:"var"`
    Name string `json:"name"`
    Leaf bool `json:"leaf"`
    Children []GetCategoryTreeDataResponseEntity `json:"children"`
}
func (g GetCategoryTreeDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}