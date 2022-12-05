package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeProductsResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    PrimaryCategory	int64	`json:"primary_category"`
    SellerSku	string	`json:"seller_sku"`
    Attributes	[]GetUnfilledAttributeAttributesResponseEntity	`json:"attributes"`
}
func (g GetUnfilledAttributeProductsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
