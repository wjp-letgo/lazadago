package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetResponseImagesResponseEntity struct {
	Url      string `json:"url"`
	HashCode string `json:"hash_code"`
}

func (g GetResponseImagesResponseEntity) String() string {
	return lib.ObjectToString(g)
}
