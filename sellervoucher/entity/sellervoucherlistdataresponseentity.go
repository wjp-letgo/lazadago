package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherListDataResponseEntity struct{
    Total	int	`json:"total"`
    Current	int	`json:"current"`
    DataList	[]SellerVoucherListDataListResponseEntity	`json:"data_list"`
    PageSize	int	`json:"page_size"`
}
func (g SellerVoucherListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
