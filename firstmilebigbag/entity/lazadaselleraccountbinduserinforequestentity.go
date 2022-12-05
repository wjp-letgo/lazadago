package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaSellerAccountBindUserInfoRequestEntity struct {
	AppUserKey string `json:"appUserKey"`
}

func (g LazadaSellerAccountBindUserInfoRequestEntity) String() string {
	return lib.ObjectToString(g)
}
