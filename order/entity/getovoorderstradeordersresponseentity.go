package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOVOOrdersTradeOrdersResponseEntity struct {
	TradeOrderId    int                                         `json:"tradeOrderId"`
	PaymentMethod   string                                      `json:"paymentMethod"`
	PaidTime        string                                      `json:"paidTime"`
	TradeOrderLines []GetOVOOrdersTradeOrderLinesResponseEntity `json:"tradeOrderLines"`
}

func (g GetOVOOrdersTradeOrdersResponseEntity) String() string {
	return lib.ObjectToString(g)
}
