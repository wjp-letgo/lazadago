package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderItemsFromBarCodeItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    ItemName	string	`json:"item_name"`
    ItemImg	string	`json:"item_img"`
    UnitFee	string	`json:"unit_fee"`
    UnitFeeCurrency	string	`json:"unit_fee_currency"`
    ActualFee	string	`json:"actual_fee"`
    ActualFeeCurrency	string	`json:"actual_fee_currency"`
}
func (g GetOrderItemsFromBarCodeItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
