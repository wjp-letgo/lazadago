package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherSelectedProductListDataResponseEntity struct{
    Total	int	`json:"total"`
    Current	int	`json:"current"`
    DataList	[]SellerVoucherSelectedProductListDataListResponseEntity	`json:"data_list"`
    PageSize	int	`json:"page_size"`
}
func (g SellerVoucherSelectedProductListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
