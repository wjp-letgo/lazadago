package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingGetPromoTierResponseEntity struct{
    Tiers	[]FreeShippingGetTiersResponseEntity	`json:"tiers"`
    DiscountType	string	`json:"discount_type"`
    DealCriteria	string	`json:"deal_criteria"`
}
func (g FreeShippingGetPromoTierResponseEntity) String() string {
    return lib.ObjectToString(g)
}
