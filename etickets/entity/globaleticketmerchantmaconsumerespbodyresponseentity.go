package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GlobalEticketMerchantMaConsumeRespBodyResponseEntity struct {
	AttributeMap GlobalEticketMerchantMaConsumeAttributeMapResponseEntity `json:"attribute_map"`
}

func (g GlobalEticketMerchantMaConsumeRespBodyResponseEntity) String() string {
	return lib.ObjectToString(g)
}

type GlobalEticketMerchantMaConsumeAttributeMapResponseEntity struct{}
