package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaBigbagCancelUserInfoRequestEntity struct {
	AppUserKey string `json:"appUserKey"`
}

func (g LazadaBigbagCancelUserInfoRequestEntity) String() string {
	return lib.ObjectToString(g)
}
