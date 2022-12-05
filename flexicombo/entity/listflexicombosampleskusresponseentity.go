package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboSampleSkusResponseEntity struct{
    ProductId	int64	`json:"product_id"`
    SkuId	int64	`json:"sku_id"`
}
func (g ListFlexiComboSampleSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
