package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboGiftSkusResponseEntity struct{
    ProductId	int64	`json:"product_id"`
    SkuId	int64	`json:"sku_id"`
}
func (g ListFlexiComboGiftSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
