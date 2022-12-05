package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type CreateFlexiComboGiftSkusRequestEntity struct {
	ProductId int64 `json:"productId"`
	SkuId     int64 `json:"skuId"`
}

func (g CreateFlexiComboGiftSkusRequestEntity) String() string {
	return lib.ObjectToString(g)
}
