package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToSOFFailedDeliveryDataResponseEntity struct{
    OrderItems	[]SetStatusToSOFFailedDeliveryOrderItemsResponseEntity	`json:"order_items"`
}
func (g SetStatusToSOFFailedDeliveryDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
