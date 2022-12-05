package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeDataResponseEntity struct{
    TotalProducts	int	`json:"total_products"`
    Products	[]GetUnfilledAttributeProductsResponseEntity	`json:"products"`
}
func (g GetUnfilledAttributeDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
