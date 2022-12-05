package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDocumentDataResponseEntity struct{
    Document	GetDocumentDocumentResponseEntity	`json:"document"`
}
func (g GetDocumentDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
