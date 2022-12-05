package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPlatformProductsDataResponseEntity struct{
    PlatformSkuName	string	`json:"platform_sku_name"`
    CreatedAt	string	`json:"created_at"`
    ExtendFields	string	`json:"extend_fields"`
    PlatformSku	string	`json:"platform_sku"`
    Source	string	`json:"source"`
    FulfillmentSkuName	string	`json:"fulfillment_sku_name"`
    Status	string	`json:"status"`
    UpdatedAt	string	`json:"updated_at"`
    Marketplace	string	`json:"marketplace"`
    FulfillmentSku	string	`json:"fulfillment_sku"`
    SellerSku	string	`json:"seller_sku"`
}
func (g GetPlatformProductsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
