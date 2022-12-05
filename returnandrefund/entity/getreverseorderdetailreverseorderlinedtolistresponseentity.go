package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderDetailReverseOrderLineDTOListResponseEntity struct{
    ReverseOrderLineId	int	`json:"reverse_order_line_id"`
    TradeOrderLineId	int	`json:"trade_order_line_id"`
    Buyer	GetReverseOrderDetailBuyerResponseEntity	`json:"buyer"`
    ReverseStatus	string	`json:"reverse_status"`
    ProductDTO	GetReverseOrderDetailProductDTOResponseEntity	`json:"productDTO"`
    IsNeedRefund	bool	`json:"is_need_refund"`
    OfcStatus	string	`json:"ofc_status"`
}
func (g GetReverseOrderDetailReverseOrderLineDTOListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
