package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateProductResult struct {
	Data      UpdateProductDataResponseEntity     `json:"data"`
	Type      string                              `json:"type"`
	Code      string                              `json:"code"`
	Message   string                              `json:"message"`
	RequestId string                              `json:"request_id"`
	Detail    []UpdateProductDetailResponseEntity `json:"detail"`
}

func (g UpdateProductResult) String() string {
	return lib.ObjectToString(g)
}

type UpdateProductDetailResponseEntity struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (g UpdateProductDetailResponseEntity) String() string {
	return lib.ObjectToString(g)
}
