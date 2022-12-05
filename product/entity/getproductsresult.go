package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetProductsResult struct {
	Data      GetProductsDataResponseEntity     `json:"data"`
	Type      string                            `json:"type"`
	Code      string                            `json:"code"`
	Message   string                            `json:"message"`
	RequestId string                            `json:"request_id"`
	Detail    []GetProductsDetailResponseEntity `json:"detail"`
}

func (g GetProductsResult) String() string {
	return lib.ObjectToString(g)
}

type GetProductsDetailResponseEntity struct{}
