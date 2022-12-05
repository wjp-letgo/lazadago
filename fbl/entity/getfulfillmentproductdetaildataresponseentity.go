package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFulfillmentProductDetailDataResponseEntity struct{
    ShelfLifeDays	int	`json:"shelf_life_days"`
    Color	string	`json:"color"`
    FulfillmentSku	string	`json:"fulfillment_sku"`
    SerialNumberFlag	bool	`json:"serial_number_flag"`
    Length	int	`json:"length"`
    CreatedAt	string	`json:"created_at"`
    OfflineShelfLive	int	`json:"offline_shelf_live"`
    Barcodes	string	`json:"barcodes"`
    NetWeight	int	`json:"net_weight"`
    AlertShelfLive	int	`json:"alert_shelf_live"`
    ShelfLifeFlag	bool	`json:"shelf_life_flag"`
    RejectShelfLive	int	`json:"reject_shelf_live"`
    UpdatedAt	string	`json:"updated_at"`
    SnSampleList	[]GetFulfillmentProductDetailSnSampleListResponseEntity	`json:"sn_sample_list"`
    Width	int	`json:"width"`
    ShipperId	string	`json:"shipper_id"`
    SerialNumberMode	string	`json:"serial_number_mode"`
    FulfillmentSkuName	string	`json:"fulfillment_sku_name"`
    GrossWeight	int	`json:"gross_weight"`
    Height	int	`json:"height"`
}
func (g GetFulfillmentProductDetailDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
