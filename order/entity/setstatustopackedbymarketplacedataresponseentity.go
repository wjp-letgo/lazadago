package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetStatusToPackedByMarketplaceDataResponseEntity struct {
	OrderItems []SetStatusToPackedByMarketplaceOrderItemsResponseEntity `json:"order_items"`
}

func (g SetStatusToPackedByMarketplaceDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
