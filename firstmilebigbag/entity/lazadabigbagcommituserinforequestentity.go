package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaBigbagCommitUserInfoRequestEntity struct {
	AppUserKey string `json:"appUserKey"`
}

func (g LazadaBigbagCommitUserInfoRequestEntity) String() string {
	return lib.ObjectToString(g)
}
