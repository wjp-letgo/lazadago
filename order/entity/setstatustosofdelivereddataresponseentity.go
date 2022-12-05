package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetStatusToSOFDeliveredDataResponseEntity struct {
	OrderItems []SetStatusToSOFDeliveredOrderItemsResponseEntity `json:"order_items"`
}

func (g SetStatusToSOFDeliveredDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
