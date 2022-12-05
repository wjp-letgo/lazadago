package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerPerformanceDataResponseEntity struct{
    SellerId	int	`json:"seller_id"`
    MainCategoryId	int	`json:"main_category_id"`
    MainCategoryName	string	`json:"main_category_name"`
    Indicators	[]GetSellerPerformanceIndicatorsResponseEntity	`json:"indicators"`
}
func (g GetSellerPerformanceDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
