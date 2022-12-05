package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetInvoiceNumberResult struct{
    Data	SetInvoiceNumberDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetInvoiceNumberDetailResponseEntity	`json:"detail"`
}
func (g SetInvoiceNumberResult) String() string {
    return lib.ObjectToString(g)
}

type SetInvoiceNumberDetailResponseEntity struct{
    
}
