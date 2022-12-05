package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UploadImageDataResponseEntity struct {
	Image UploadImageImageResponseEntity `json:"image"`
}

func (g UploadImageDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
