package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ReverseOrderReturnUpdateReasonInfoResponseEntity struct {
	ReasonId   int    `json:"reason_id"`
	ReasonName string `json:"reason_name"`
}

func (g ReverseOrderReturnUpdateReasonInfoResponseEntity) String() string {
	return lib.ObjectToString(g)
}
