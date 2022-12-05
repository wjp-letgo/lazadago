package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetLazadaBigbagPDFLableResult struct {
	Result    GetLazadaBigbagPDFLableResultResponseEntity   `json:"result"`
	Type      string                                        `json:"type"`
	Code      string                                        `json:"code"`
	Message   string                                        `json:"message"`
	RequestId string                                        `json:"request_id"`
	Detail    []GetLazadaBigbagPDFLableDetailResponseEntity `json:"detail"`
}

func (g GetLazadaBigbagPDFLableResult) String() string {
	return lib.ObjectToString(g)
}

type GetLazadaBigbagPDFLableDetailResponseEntity struct{}
