package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingSelectedProductListDataListResponseEntity struct{
    ProductId	int64	`json:"product_id"`
    SkuIds	[]int64	`json:"sku_ids"`
}
func (g FreeShippingSelectedProductListDataListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
