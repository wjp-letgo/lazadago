package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetQcStatusDataResponseEntity struct{
    SellerSku	string	`json:"seller_sku"`
    Status	string	`json:"status"`
    DataChanged	string	`json:"data_changed"`
    Reason	string	`json:"reason"`
}
func (g GetQcStatusDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
