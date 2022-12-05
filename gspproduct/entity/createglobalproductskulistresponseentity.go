package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateGlobalProductSkuListResponseEntity struct{
    SellerSku	string	`json:"seller_sku"`
    ShopSku	string	`json:"shop_sku"`
    SkuId	int64	`json:"sku_id"`
}
func (g CreateGlobalProductSkuListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
