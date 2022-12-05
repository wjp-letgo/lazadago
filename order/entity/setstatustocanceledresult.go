package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToCanceledResult struct{
    Success	bool	`json:"success"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetStatusToCanceledDetailResponseEntity	`json:"detail"`
}
func (g SetStatusToCanceledResult) String() string {
    return lib.ObjectToString(g)
}

type SetStatusToCanceledDetailResponseEntity struct{
    
}
