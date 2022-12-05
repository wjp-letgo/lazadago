package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOutboundOrderDetailSkusResponseEntity struct{
    SellerSku	string	`json:"seller_sku"`
    ItemOutbounded	int	`json:"item_outbounded"`
    SerialNumberFlag	bool	`json:"serial_number_flag"`
    SkuStatus	string	`json:"sku_status"`
    RequestedQuantity	int	`json:"requested_quantity"`
    FulfillmentSkuName	string	`json:"fulfillment_sku_name"`
    ShelfLifeFlag	bool	`json:"shelf_life_flag"`
    Barcodes	string	`json:"barcodes"`
    FulfillmentSku	string	`json:"fulfillment_sku"`
    Comments	string	`json:"comments"`
}
func (g GetOutboundOrderDetailSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
