package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAwbDocumentHtmlDocumentResponseEntity struct{
    File	string	`json:"file"`
    MimeType	string	`json:"mime_type"`
    DocumentType	string	`json:"document_type"`
}
func (g GetAwbDocumentHtmlDocumentResponseEntity) String() string {
    return lib.ObjectToString(g)
}
