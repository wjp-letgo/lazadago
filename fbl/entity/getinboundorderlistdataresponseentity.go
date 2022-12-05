package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInboundOrderListDataResponseEntity struct{
    SkuApproved	int	`json:"sku_approved"`
    InboundTime	string	`json:"inbound_time"`
    IoNumber	string	`json:"io_number"`
    EstimateTime	string	`json:"estimate_time"`
    Marketplace	string	`json:"marketplace"`
    ItemInboundedGood	int	`json:"item_inbounded_good"`
    DeliveryType	string	`json:"delivery_type"`
    ItemRequested	int	`json:"item_requested"`
    CreatedAt	string	`json:"created_at"`
    SkuInbounded	int	`json:"sku_inbounded"`
    SkuRequested	int	`json:"sku_requested"`
    InboundWarehouse	string	`json:"inbound_warehouse"`
    ReferenceNumber	string	`json:"reference_number"`
    ItemInboundedDefective	int	`json:"item_inbounded_defective"`
    UpdatedAt	string	`json:"updated_at"`
    Status	string	`json:"status"`
    ShopName	string	`json:"shop_name"`
    IoType	string	`json:"io_type"`
}
func (g GetInboundOrderListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
