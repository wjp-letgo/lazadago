package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetStatusToSOFFailedDeliveryDataResponseEntity struct {
	OrderItems []SetStatusToSOFFailedDeliveryOrderItemsResponseEntity `json:"order_items"`
}

func (g SetStatusToSOFFailedDeliveryDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
