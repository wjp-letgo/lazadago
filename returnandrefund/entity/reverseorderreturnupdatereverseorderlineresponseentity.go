package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ReverseOrderReturnUpdateReverseOrderLineResponseEntity struct{
    ReverseOrderLineId	int	`json:"reverse_order_line_id"`
    ReasonSource	string	`json:"reason_source"`
    ReasonType	string	`json:"reason_type"`
    ReasonId	int	`json:"reason_id"`
    ReasonName	string	`json:"reason_name"`
    ReasonDesc	string	`json:"reason_desc"`
    RefundAmount	int	`json:"refund_amount"`
    IsCancel	bool	`json:"is_cancel"`
    OrderId	int64	`json:"order_id"`
    SellerSku	string	`json:"seller_sku"`
    PaidPrice	float32	`json:"paid_price"`
    ApplyReason	string	`json:"apply_reason"`
    OrderLineId	int	`json:"order_line_id"`
}
func (g ReverseOrderReturnUpdateReverseOrderLineResponseEntity) String() string {
    return lib.ObjectToString(g)
}
