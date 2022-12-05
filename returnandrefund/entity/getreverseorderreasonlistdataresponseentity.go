package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderReasonListDataResponseEntity struct{
    ReasonId	int	`json:"reason_id"`
    MutiLanguageText	string	`json:"muti_language_text"`
    Text	string	`json:"text"`
}
func (g GetReverseOrderReasonListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
