package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOrdersResult struct {
	Data      GetOrdersDataResponseEntity     `json:"data"`
	Type      string                          `json:"type"`
	Code      string                          `json:"code"`
	Message   string                          `json:"message"`
	RequestId string                          `json:"request_id"`
	Detail    []GetOrdersDetailResponseEntity `json:"detail"`
}

func (g GetOrdersResult) String() string {
	return lib.ObjectToString(g)
}

type GetOrdersDetailResponseEntity struct {
}
