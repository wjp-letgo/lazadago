package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type MigrateImageDataResponseEntity struct {
	Image MigrateImageImageResponseEntity `json:"image"`
}

func (g MigrateImageDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
