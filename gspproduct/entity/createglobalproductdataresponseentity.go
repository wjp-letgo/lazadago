package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateGlobalProductDataResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    SkuList	[]CreateGlobalProductSkuListResponseEntity	`json:"sku_list"`
}
func (g CreateGlobalProductDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
