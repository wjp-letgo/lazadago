package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetReverseOrdersForSellerBuyerResponseEntity struct {
	BuyerId int64 `json:"buyer_id"`
}

func (g GetReverseOrdersForSellerBuyerResponseEntity) String() string {
	return lib.ObjectToString(g)
}
