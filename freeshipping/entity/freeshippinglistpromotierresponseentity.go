package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingListPromoTierResponseEntity struct{
    Tiers	[]FreeShippingListTiersResponseEntity	`json:"tiers"`
    DiscountType	string	`json:"discount_type"`
    DealCriteria	string	`json:"deal_criteria"`
}
func (g FreeShippingListPromoTierResponseEntity) String() string {
    return lib.ObjectToString(g)
}
