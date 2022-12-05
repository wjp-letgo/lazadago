package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaBigbagCollectionPointsResultResponseEntity struct {
	Data      LazadaBigbagCollectionPointsDataResponseEntity `json:"data"`
	Success   bool                                           `json:"success"`
	ErrorCode string                                         `json:"errorCode"`
	ErroMsg   string                                         `json:"erroMsg"`
}

func (g LazadaBigbagCollectionPointsResultResponseEntity) String() string {
	return lib.ObjectToString(g)
}
