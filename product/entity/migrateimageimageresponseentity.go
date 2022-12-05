package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type MigrateImageImageResponseEntity struct {
	Url      string `json:"url"`
	HashCode string `json:"hash_code"`
}

func (g MigrateImageImageResponseEntity) String() string {
	return lib.ObjectToString(g)
}
