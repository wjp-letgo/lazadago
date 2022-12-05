package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrdersForSellerReverseOrderLinesResponseEntity struct{
    ReverseOrderLineId	int	`json:"reverse_order_line_id"`
    TradeOrderLineId	int	`json:"trade_order_line_id"`
    ReverseStatus	string	`json:"reverse_status"`
    IsNeedRefund	string	`json:"is_need_refund"`
    OfcStatus	string	`json:"ofc_status"`
    Product	GetReverseOrdersForSellerProductResponseEntity	`json:"product"`
    Buyer	GetReverseOrdersForSellerBuyerResponseEntity	`json:"buyer"`
}
func (g GetReverseOrdersForSellerReverseOrderLinesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
