package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutStatusDataResponseEntity struct{
    ClosingBalance	string	`json:"closing_balance"`
    GuaranteeDeposit	string	`json:"guarantee_deposit"`
    Payout	string	`json:"payout"`
    Paid	string	`json:"paid"`
    StatementNumber	string	`json:"statement_number"`
    CreatedAt	string	`json:"created_at"`
    UpdatedAt	string	`json:"updated_at"`
    OpeningBalance	string	`json:"opening_balance"`
    ItemRevenue	string	`json:"item_revenue"`
    ShipmentFee	string	`json:"shipment_fee"`
    ShipmentFeeCredit	string	`json:"shipment_fee_credit"`
    OtherRevenueTotal	string	`json:"other_revenue_total"`
    FeesTotal	string	`json:"fees_total"`
    Subtotal1	string	`json:"subtotal1"`
    Refunds	string	`json:"refunds"`
    FeesOnRefundsTotal	string	`json:"fees_on_refunds_total"`
    Subtotal2	string	`json:"subtotal2"`
}
func (g GetPayoutStatusDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
