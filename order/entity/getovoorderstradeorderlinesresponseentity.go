package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOVOOrdersTradeOrderLinesResponseEntity struct{
    DeliveredTime	string	`json:"deliveredTime"`
    TradeOrderLineId	int	`json:"tradeOrderLineId"`
    DeliveryStatus	string	`json:"deliveryStatus"`
    ReverseStatus	string	`json:"reverseStatus"`
}
func (g GetOVOOrdersTradeOrderLinesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
