package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaAvailableRespBodyResponseEntity struct{
    AttributeMap	GlobalEticketMerchantMaAvailableAttributeMapResponseEntity	`json:"attribute_map"`
}
func (g GlobalEticketMerchantMaAvailableRespBodyResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaAvailableAttributeMapResponseEntity struct{}