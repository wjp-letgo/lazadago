package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToReadyToShipDataResponseEntity struct{
    OrderItems	[]SetStatusToReadyToShipOrderItemsResponseEntity	`json:"order_items"`
}
func (g SetStatusToReadyToShipDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
