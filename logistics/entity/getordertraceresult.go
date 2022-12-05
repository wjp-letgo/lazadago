package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOrderTraceResult struct {
	Result    GetOrderTraceResultResponseEntity   `json:"result"`
	Type      string                              `json:"type"`
	Code      string                              `json:"code"`
	Message   string                              `json:"message"`
	RequestId string                              `json:"request_id"`
	Detail    []GetOrderTraceDetailResponseEntity `json:"detail"`
}

func (g GetOrderTraceResult) String() string {
	return lib.ObjectToString(g)
}

type GetOrderTraceDetailResponseEntity struct{}
