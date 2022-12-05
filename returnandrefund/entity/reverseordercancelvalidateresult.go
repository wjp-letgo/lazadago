package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ReverseOrderCancelValidateResult struct {
	Data      ReverseOrderCancelValidateDataResponseEntity     `json:"data"`
	Type      string                                           `json:"type"`
	Code      string                                           `json:"code"`
	Message   string                                           `json:"message"`
	RequestId string                                           `json:"request_id"`
	Detail    []ReverseOrderCancelValidateDetailResponseEntity `json:"detail"`
}

func (g ReverseOrderCancelValidateResult) String() string {
	return lib.ObjectToString(g)
}

type ReverseOrderCancelValidateDetailResponseEntity struct{}
