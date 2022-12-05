package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetStatusToReadyToShipDataResponseEntity struct {
	OrderItems []SetStatusToReadyToShipOrderItemsResponseEntity `json:"order_items"`
}

func (g SetStatusToReadyToShipDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
