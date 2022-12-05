package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOutboundOrderListDataResponseEntity struct{
    SkuApproved	int	`json:"sku_approved"`
    OutboundTime	string	`json:"outbound_time"`
    OoNumber	string	`json:"oo_number"`
    EstimateTime	string	`json:"estimate_time"`
    Marketplace	string	`json:"marketplace"`
    DeliveryType	string	`json:"delivery_type"`
    ItemRequested	int	`json:"item_requested"`
    CreatedAt	string	`json:"created_at"`
    SkuOutbounded	int	`json:"sku_outbounded"`
    SkuRequested	int	`json:"sku_requested"`
    OutboundWarehouse	string	`json:"outbound_warehouse"`
    ReferenceNumber	string	`json:"reference_number"`
    ItemOutbounded	int	`json:"item_outbounded"`
    UpdatedAt	string	`json:"updated_at"`
    Status	string	`json:"status"`
    ShopName	string	`json:"shop_name"`
    CreatedBy	string	`json:"created_by"`
    OutboundReason	string	`json:"outbound_reason"`
    FulfillmentOrderNumber	string	`json:"fulfillment_order_number"`
}
func (g GetOutboundOrderListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
