package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadSellerBagNumberDataResponseEntity struct{
    ErpBagNumber	string	`json:"erp_bag_number"`
    Country	string	`json:"country"`
    ResponseDate	string	`json:"response_date"`
    LzdBagNumber	string	`json:"lzd_bag_number"`
    GlobalCollection	string	`json:"global_collection"`
    SellerId	string	`json:"seller_id"`
    LabelUrl	string	`json:"label_url"`
}
func (g UploadSellerBagNumberDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
