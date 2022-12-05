package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateProductDataResponseEntity struct {
	Variation UpdateProductVariationResponseEntity `json:"variation"`
}

func (g UpdateProductDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
