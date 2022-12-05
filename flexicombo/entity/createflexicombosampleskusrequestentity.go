package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type CreateFlexiComboSampleSkusRequestEntity struct {
	ProductId int64 `json:"productId"`
	SkuId     int64 `json:"skuId"`
}

func (g CreateFlexiComboSampleSkusRequestEntity) String() string {
	return lib.ObjectToString(g)
}
