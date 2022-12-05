package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GlobalEticketMerchantMaAvailableRespBodyResponseEntity struct {
	AttributeMap GlobalEticketMerchantMaAvailableAttributeMapResponseEntity `json:"attribute_map"`
}

func (g GlobalEticketMerchantMaAvailableRespBodyResponseEntity) String() string {
	return lib.ObjectToString(g)
}

type GlobalEticketMerchantMaAvailableAttributeMapResponseEntity struct{}
