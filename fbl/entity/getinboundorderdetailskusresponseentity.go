package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInboundOrderDetailSkusResponseEntity struct{
    SellerSku	string	`json:"seller_sku"`
    ItemInboundedGood	int	`json:"item_inbounded_good"`
    SerialNumberFlag	bool	`json:"serial_number_flag"`
    SkuStatus	string	`json:"sku_status"`
    ItemInboundedDefective	int	`json:"item_inbounded_defective"`
    FulfillmentSkuName	string	`json:"fulfillment_sku_name"`
    RequestedQuantity	int	`json:"requested_quantity"`
    ShelfLifeFlag	bool	`json:"shelf_life_flag"`
    Barcodes	string	`json:"barcodes"`
    FulfillmentSku	string	`json:"fulfillment_sku"`
    Comments	string	`json:"comments"`
}
func (g GetInboundOrderDetailSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
