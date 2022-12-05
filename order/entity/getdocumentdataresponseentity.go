package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetDocumentDataResponseEntity struct {
	Document GetDocumentDocumentResponseEntity `json:"document"`
}

func (g GetDocumentDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
