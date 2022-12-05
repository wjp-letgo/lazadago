package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SellerVoucherCreateResult struct {
	Data      int                                       `json:"data"`
	Success   bool                                      `json:"success"`
	ErrorCode int                                       `json:"error_code"`
	ErrorMsg  string                                    `json:"error_msg"`
	Type      string                                    `json:"type"`
	Code      string                                    `json:"code"`
	Message   string                                    `json:"message"`
	RequestId string                                    `json:"request_id"`
	Detail    []SellerVoucherCreateDetailResponseEntity `json:"detail"`
}

func (g SellerVoucherCreateResult) String() string {
	return lib.ObjectToString(g)
}

type SellerVoucherCreateDetailResponseEntity struct{}
