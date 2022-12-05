package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategorySuggestionCategorySuggestionsResponseEntity struct{
    CategoryId	int	`json:"categoryId"`
    CategoryName	string	`json:"categoryName"`
    CategoryPath	string	`json:"categoryPath"`
}
func (g GetCategorySuggestionCategorySuggestionsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
