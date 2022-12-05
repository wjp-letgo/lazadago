package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetNewBagNumberResult struct {
	Result    GetNewBagNumberResultResponseEntity   `json:"result"`
	Type      string                                `json:"type"`
	Code      string                                `json:"code"`
	Message   string                                `json:"message"`
	RequestId string                                `json:"request_id"`
	Detail    []GetNewBagNumberDetailResponseEntity `json:"detail"`
}

func (g GetNewBagNumberResult) String() string {
	return lib.ObjectToString(g)
}

type GetNewBagNumberDetailResponseEntity struct{}
