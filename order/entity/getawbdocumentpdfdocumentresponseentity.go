package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetAwbDocumentPDFDocumentResponseEntity struct {
	File         string `json:"file"`
	MimeType     string `json:"mime_type"`
	DocumentType string `json:"document_type"`
}

func (g GetAwbDocumentPDFDocumentResponseEntity) String() string {
	return lib.ObjectToString(g)
}
