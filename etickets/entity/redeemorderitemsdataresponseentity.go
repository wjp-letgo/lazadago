package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RedeemOrderItemsDataResponseEntity struct{
    OuterId	string	`json:"outer_id"`
    SerialNum	string	`json:"serial_num"`
    LeftNum	int	`json:"left_num"`
}
func (g RedeemOrderItemsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
