package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaBigbagUpdateUserInfoRequestEntity struct {
	AppUserKey string `json:"appUserKey"`
}

func (g LazadaBigbagUpdateUserInfoRequestEntity) String() string {
	return lib.ObjectToString(g)
}
