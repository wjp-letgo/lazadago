package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategorySuggestionDataResponseEntity struct{
    CategorySuggestions	[]GetCategorySuggestionCategorySuggestionsResponseEntity	`json:"categorySuggestions"`
}
func (g GetCategorySuggestionDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
