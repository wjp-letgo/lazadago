package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetResponseDataResponseEntity struct {
	Images []GetResponseImagesResponseEntity `json:"images"`
	Errors []GetResponseErrorsResponseEntity `json:"errors"`
}

func (g GetResponseDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
