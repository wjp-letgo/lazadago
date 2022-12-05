package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInboundOrderDetailDataResponseEntity struct{
    InboundTime	string	`json:"inbound_time"`
    Skus	[]GetInboundOrderDetailSkusResponseEntity	`json:"skus"`
    Comments	string	`json:"comments"`
    IoNumber	string	`json:"io_number"`
    EstimateTime	string	`json:"estimate_time"`
    Marketplace	string	`json:"marketplace"`
    DeliveryType	string	`json:"delivery_type"`
    CreatedAt	string	`json:"created_at"`
    InboundWarehouse	string	`json:"inbound_warehouse"`
    ReferenceNumber	string	`json:"reference_number"`
    UpdatedAt	string	`json:"updated_at"`
    IoStatus	string	`json:"io_status"`
    ShopName	string	`json:"shop_name"`
    IoType	string	`json:"io_type"`
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
func (g GetInboundOrderDetailDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
