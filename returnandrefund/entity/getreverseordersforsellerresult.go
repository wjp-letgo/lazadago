package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetReverseOrdersForSellerResult struct {
	Result    GetReverseOrdersForSellerResultResponseEntity   `json:"result"`
	Type      string                                          `json:"type"`
	Code      string                                          `json:"code"`
	Message   string                                          `json:"message"`
	RequestId string                                          `json:"request_id"`
	Detail    []GetReverseOrdersForSellerDetailResponseEntity `json:"detail"`
}

func (g GetReverseOrdersForSellerResult) String() string {
	return lib.ObjectToString(g)
}

type GetReverseOrdersForSellerDetailResponseEntity struct{}
