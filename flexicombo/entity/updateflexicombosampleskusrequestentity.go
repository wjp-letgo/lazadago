package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateFlexiComboSampleSkusRequestEntity struct{
    ProductId	int64	`json:"productId"`
    SkuId	int64	`json:"skuId"`
}
func (g UpdateFlexiComboSampleSkusRequestEntity) String() string {
    return lib.ObjectToString(g)
}
