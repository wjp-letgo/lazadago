package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetInvoiceNumberDataResponseEntity struct {
	OrderItemId   int64  `json:"order_item_id"`
	InvoiceNumber string `json:"invoice_number"`
}

func (g SetInvoiceNumberDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
