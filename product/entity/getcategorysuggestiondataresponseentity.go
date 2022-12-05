package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetCategorySuggestionDataResponseEntity struct {
	CategorySuggestions []GetCategorySuggestionCategorySuggestionsResponseEntity `json:"categorySuggestions"`
}

func (g GetCategorySuggestionDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
