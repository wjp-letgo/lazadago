package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ReverseOrderCancelValidateDataResponseEntity struct{
    TipContent	string	`json:"tip_content"`
    TipType	string	`json:"tip_type"`
    ReasonOptions	[]ReverseOrderCancelValidateReasonOptionsResponseEntity	`json:"reason_options"`
}
func (g ReverseOrderCancelValidateDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
