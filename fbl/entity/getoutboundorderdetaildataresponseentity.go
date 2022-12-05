package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOutboundOrderDetailDataResponseEntity struct{
    OutboundTime	string	`json:"outbound_time"`
    Comments	string	`json:"comments"`
    Skus	[]GetOutboundOrderDetailSkusResponseEntity	`json:"skus"`
    EstimateTime	string	`json:"estimate_time"`
    Marketplace	string	`json:"marketplace"`
    OutboundWarehouse	string	`json:"outbound_warehouse"`
    DeliveryType	string	`json:"delivery_type"`
    CreatedAt	string	`json:"created_at"`
    ReferenceNumber	string	`json:"reference_number"`
    ItemOutbounded	int	`json:"item_outbounded"`
    OutboundOrderNo	string	`json:"outbound_order_no"`
    UpdatedAt	string	`json:"updated_at"`
    Status	string	`json:"status"`
    ShopName	string	`json:"shop_name"`
    CreatedBy	string	`json:"created_by"`
    OutboundReason	string	`json:"outbound_reason"`
    InventoryType	string	`json:"inventory_type"`
    WarehouseName	string	`json:"warehouse_name"`
    WarehouseAddress	string	`json:"warehouse_address"`
    SellerWarehouseName	string	`json:"seller_warehouse_name"`
    SellerCity	string	`json:"seller_city"`
    SellerAddress	string	`json:"seller_address"`
    SellerPostcode	string	`json:"seller_postcode"`
    SellerCountry	string	`json:"seller_country"`
    SellerContact	string	`json:"seller_contact"`
    SellerMobile	string	`json:"seller_mobile"`
    FulfillmentOrderNumber	string	`json:"fulfillment_order_number"`
}
func (g GetOutboundOrderDetailDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
