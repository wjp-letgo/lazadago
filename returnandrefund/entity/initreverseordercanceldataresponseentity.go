package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type InitReverseOrderCancelDataResponseEntity struct{
    TipContent	string	`json:"tip_content"`
    TipType	string	`json:"tip_type"`
}
func (g InitReverseOrderCancelDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
