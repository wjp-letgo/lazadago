package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetFailureReasonsDataResponseEntity struct {
	ReasonId int    `json:"reason_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
}

func (g GetFailureReasonsDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
