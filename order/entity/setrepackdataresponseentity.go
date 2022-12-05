package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetRepackDataResponseEntity struct {
	PackageId string `json:"package_id"`
}

func (g SetRepackDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
