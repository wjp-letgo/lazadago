package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAwbDocumentPDFResult struct{
    Data	GetAwbDocumentPDFDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetAwbDocumentPDFDetailResponseEntity	`json:"detail"`
}
func (g GetAwbDocumentPDFResult) String() string {
    return lib.ObjectToString(g)
}

type GetAwbDocumentPDFDetailResponseEntity struct{

}