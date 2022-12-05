package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaConsumeRespBodyResponseEntity struct{
    AttributeMap	GlobalEticketMerchantMaConsumeAttributeMapResponseEntity	`json:"attribute_map"`
}
func (g GlobalEticketMerchantMaConsumeRespBodyResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaConsumeAttributeMapResponseEntity struct{}