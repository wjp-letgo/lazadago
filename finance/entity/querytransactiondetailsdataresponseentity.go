package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryTransactionDetailsDataResponseEntity struct{
    FeeType	string	`json:"fee_type"`
    Details	string	`json:"details"`
    SellerSku	string	`json:"seller_sku"`
    LazadaSku	string	`json:"lazada_sku"`
    Amount	string	`json:"amount"`
    VATInAmount	string	`json:"VAT_in_amount"`
    WHTAmount	string	`json:"WHT_amount"`
    WHTIncludedInAmount	string	`json:"WHT_included_in_amount"`
    Statement	string	`json:"statement"`
    PaidStatus	string	`json:"paid_status"`
    OrderNo	string	`json:"order_no"`
    OrderItemNo	string	`json:"orderItem_no"`
    OrderItemStatus	string	`json:"orderItem_status"`
    ShippingProvider	string	`json:"shipping_provider"`
    ShippingSpeed	string	`json:"shipping_speed"`
    ShipmentType	string	`json:"shipment_type"`
    Reference	string	`json:"reference"`
    Comment	string	`json:"comment"`
    PaymentRefId	string	`json:"payment_ref_id"`
    FeeName	string	`json:"fee_name"`
    TransactionDate	string	`json:"transaction_date"`
    TransactionType	string	`json:"transaction_type"`
    TransactionNumber	string	`json:"transaction_number"`
}
func (g QueryTransactionDetailsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
