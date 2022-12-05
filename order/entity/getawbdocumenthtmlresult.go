package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAwbDocumentHtmlResult struct{
    Data	GetAwbDocumentHtmlDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetAwbDocumentHtmlDetailResponseEntity	`json:"detail"`
}
func (g GetAwbDocumentHtmlResult) String() string {
    return lib.ObjectToString(g)
}
type GetAwbDocumentHtmlDetailResponseEntity struct{
    
}