package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetStatusToSOFFailedDeliveryOrderItemsResponseEntity struct {
	OrderItemId         int64  `json:"order_item_id"`
	PurchaseOrderId     int64  `json:"purchase_order_id"`
	PurchaseOrderNumber string `json:"purchase_order_number"`
}

func (g SetStatusToSOFFailedDeliveryOrderItemsResponseEntity) String() string {
	return lib.ObjectToString(g)
}
