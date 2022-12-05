package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDocumentResult struct{
    Data	GetDocumentDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetDocumentDetailResponseEntity	`json:"detail"`
}
func (g GetDocumentResult) String() string {
    return lib.ObjectToString(g)
}
type GetDocumentDetailResponseEntity struct{
    
}