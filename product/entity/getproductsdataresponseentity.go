package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductsDataResponseEntity struct{
    TotalProducts	int	`json:"total_products"`
    Products	[]GetProductsProductsResponseEntity	`json:"products"`
}
func (g GetProductsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
