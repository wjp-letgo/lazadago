package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMultipleOrderItemsResult struct {
	Data      []GetMultipleOrderItemsDataResponseEntity   `json:"data"`
	Type      string                                      `json:"type"`
	Code      string                                      `json:"code"`
	Message   string                                      `json:"message"`
	RequestId string                                      `json:"request_id"`
	Detail    []GetMultipleOrderItemsDetailResponseEntity `json:"detail"`
}

func (g GetMultipleOrderItemsResult) String() string {
	return lib.ObjectToString(g)
}

type GetMultipleOrderItemsDetailResponseEntity struct {
}
