package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderItemsDataResponseEntity struct{
    PurchaseOrderNumber	string	`json:"purchase_order_number"`
    Name	string	`json:"name"`
    ProductMainImage	string	`json:"product_main_image"`
    ItemPrice	float32	`json:"item_price"`
    TaxAmount	float32	`json:"tax_amount"`
    Status	string	`json:"status"`
    CancelReturnInitiator	string	`json:"cancel_return_initiator"`
    VoucherPlatform	float32	`json:"voucher_platform"`
    VoucherSeller	string	`json:"voucher_seller"`
    OrderType	string	`json:"order_type"`
    StagePayStatus	string	`json:"stage_pay_status"`
    WarehouseCode	string	`json:"warehouse_code"`
    VoucherSellerLpi	float32	`json:"voucher_seller_lpi"`
    VoucherPlatformLpi	float32	`json:"voucher_platform_lpi"`
    BuyerId	int64	`json:"buyer_id"`
    ShippingFeeOriginal	float32	`json:"shipping_fee_original"`
    ShippingFeeDiscountSeller	float32	`json:"shipping_fee_discount_seller"`
    ShippingFeeDiscountPlatform	float32	`json:"shipping_fee_discount_platform"`
    VoucherCodeSeller	string	`json:"voucher_code_seller"`
    VoucherCodePlatform	string	`json:"voucher_code_platform"`
    DeliveryOptionSof	int	`json:"delivery_option_sof"`
    IsFbl	int	`json:"is_fbl"`
    IsReroute	int	`json:"is_reroute"`
    Reason	string	`json:"reason"`
    DigitalDeliveryInfo	string	`json:"digital_delivery_info"`
    PromisedShippingTime	string	`json:"promised_shipping_time"`
    OrderId	int64	`json:"order_id"`
    VoucherAmount	float32	`json:"voucher_amount"`
    ReturnStatus	string	`json:"return_status"`
    ShippingType	string	`json:"shipping_type"`
    ShipmentProvider	string	`json:"shipment_provider"`
    Variation	string	`json:"variation"`
    CreatedAt	string	`json:"created_at"`
    InvoiceNumber	string	`json:"invoice_number"`
    ShippingAmount	float32	`json:"shipping_amount"`
    Currency	string	`json:"currency"`
    OrderFlag	string	`json:"order_flag"`
    ShopId	string	`json:"shop_id"`
    SlaTimeStamp	string	`json:"sla_time_stamp"`
    Sku	string	`json:"sku"`
    VoucherCode	string	`json:"voucher_code"`
    WalletCredits	float32	`json:"wallet_credits"`
    UpdatedAt	string	`json:"updated_at"`
    IsDigital	int	`json:"is_digital"`
    TrackingCodePre	string	`json:"tracking_code_pre"`
    OrderItemId	int64	`json:"order_item_id"`
    PackageId	string	`json:"package_id"`
    TrackingCode	string	`json:"tracking_code"`
    ShippingServiceCost	float32	`json:"shipping_service_cost"`
    ExtraAttributes	string	`json:"extra_attributes"`
    PaidPrice	float32	`json:"paid_price"`
    ShippingProviderType	string	`json:"shipping_provider_type"`
    ProductDetailUrl	string	`json:"product_detail_url"`
    ShopSku	string	`json:"shop_sku"`
    ReasonDetail	string	`json:"reason_detail"`
    PurchaseOrderId	string	`json:"purchase_order_id"`
}
func (g GetOrderItemsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
