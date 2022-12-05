package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingSelectedProductListDataResponseEntity struct{
    Total	int	`json:"total"`
    Current	int	`json:"current"`
    DataList	[]FreeShippingSelectedProductListDataListResponseEntity	`json:"data_list"`
    PageSize	int	`json:"page_size"`
}
func (g FreeShippingSelectedProductListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
