package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ReverseOrderReturnUpdateDataResponseEntity struct{
    ReverseOrderLine	[]ReverseOrderReturnUpdateReverseOrderLineResponseEntity	`json:"reverse_order_line"`
    ReverseOrderId	int64	`json:"reverse_order_id"`
    ReasonInfo	[]ReverseOrderReturnUpdateReasonInfoResponseEntity	`json:"reason_info"`
    TotalRefund	string	`json:"total_refund"`
}
func (g ReverseOrderReturnUpdateDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
