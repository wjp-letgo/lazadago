package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrdersForSellerItemsResponseEntity struct{
    ReverseOrderId	int64	`json:"reverse_order_id"`
    TradeOrderId	int64	`json:"trade_order_id"`
    RequestType	string	`json:"request_type"`
    IsRtm	bool	`json:"is_rtm"`
    ShippingType	string	`json:"shipping_type"`
    ReverseOrderLines	[]GetReverseOrdersForSellerReverseOrderLinesResponseEntity	`json:"reverse_order_lines"`
}
func (g GetReverseOrdersForSellerItemsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
