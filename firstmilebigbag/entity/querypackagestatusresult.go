package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type QueryPackageStatusResult struct {
	Result    QueryPackageStatusResultResponseEntity   `json:"result"`
	Type      string                                   `json:"type"`
	Code      string                                   `json:"code"`
	Message   string                                   `json:"message"`
	RequestId string                                   `json:"request_id"`
	Detail    []QueryPackageStatusDetailResponseEntity `json:"detail"`
}

func (g QueryPackageStatusResult) String() string {
	return lib.ObjectToString(g)
}

type QueryPackageStatusDetailResponseEntity struct{}
