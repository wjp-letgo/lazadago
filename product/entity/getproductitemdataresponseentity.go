package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductItemDataResponseEntity struct{
    Variation	GetProductItemVariationResponseEntity	`json:"variation"`
    PrimaryCategory	int64	`json:"primary_category"`
    Attributes	GetProductItemAttributesResponseEntity	`json:"attributes"`
    Skus	[]GetProductItemSkusResponseEntity	`json:"skus"`
    ItemId	int64	`json:"item_id"`
    CreatedTime	string	`json:"created_time"`
    UpdatedTime	string	`json:"updated_time"`
    Images	[]string	`json:"images"`
}
func (g GetProductItemDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
//GetProductItemAttributesResponseEntity
type GetProductItemAttributesResponseEntity struct{
    ShortDescription string `json:"short_description"`
    Name string `json:"name"`
    Description string `json:"description"`
    WarrantyType string `json:"warranty_type"`
    Brand string `json:"brand"`
    Warranty string `json:"warranty"`
    Hazmat string `json:"Hazmat"`
    Source string `json:"source"`
    Model string `json:"model"`
    ColorFamily string `json:"color_family"`
}

func (g GetProductItemAttributesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
//GetProductItemSkusResponseEntity
type GetProductItemSkusResponseEntity struct{
    Status string `json:"Status"`
    SkuId int64 `json:"SkuId"`
    Quantity int `json:"quantity"`
    ProductWeight string `json:"product_weight"`
    Images []string `json:"Images"`
    SellerSku string `json:"SellerSku"`
    ShopSku string `json:"ShopSku"`
    Url string `json:"Url"`
    PackageWidth string `json:"package_width"`
    SpecialToTime string `json:"special_to_time"`
    SpecialFromTime string `json:"special_from_time"`
    PackageHeight string `json:"package_height"`
    SpecialPrice float32 `json:"special_price"`
    Price float32 `json:"price"`
    PackageLength string `json:"package_length"`
    PackageWeight string `json:"package_weight"`
    Available int `json:"Available"`
    SpecialoDate string `json:"special_to_date"`
    ColorFamily string `json:"color_family"`
    Size string `json:"size"`
    ColorThumbnail string `json:"color_thumbnail"`
    ColorText string `json:"color_text"`
    SpecialTimeFormat string `json:"special_time_format"`
    PackageContent string `json:"package_content"`
    MultiWarehouseInventories []MultiWarehouseInventories `json:"multiWarehouseInventories"`
}

func (g GetProductItemSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}
//MultiWarehouseInventories
type MultiWarehouseInventories struct{
    WarehouseCode string `json:"warehouseCode"`
    Quantity int `json:"quantity"`
    OccupyQuantity int `json:"occupyQuantity"`
    TotalQuantity int `json:"totalQuantity"`
    WithholdQuantity int `json:"withholdQuantity"`
    SellableQuantity int `json:"sellableQuantity"`
}
func (g MultiWarehouseInventories) String() string {
    return lib.ObjectToString(g)
}