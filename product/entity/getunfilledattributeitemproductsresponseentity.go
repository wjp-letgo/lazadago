package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeItemProductsResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    PrimaryCategory	int64	`json:"primary_category"`
    Attributes	[]GetUnfilledAttributeItemAttributesResponseEntity	`json:"attributes"`
    SellerSkuId	string	`json:"seller_sku_id"`
}
func (g GetUnfilledAttributeItemProductsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
