package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadLazadaBagNumberDataResponseEntity struct{
    ErpBagNumber	string	`json:"erp_bag_number"`
    Country	string	`json:"country"`
    SellerName	string	`json:"seller_name"`
    ResponseDate	string	`json:"response_date"`
    LzdBagNumber	string	`json:"lzd_bag_number"`
    GlobalCollection	string	`json:"global_collection"`
    SellerId	string	`json:"seller_id"`
    LabelUrl	string	`json:"label_url"`
}
func (g UploadLazadaBagNumberDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
