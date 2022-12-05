package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerMetricsByIdDataResponseEntity struct{
    MainCategoryName	string	`json:"main_category_name"`
    SellerId	int	`json:"seller_id"`
    ResponseRate	string	`json:"response_rate"`
    ResponseTime	string	`json:"response_time"`
    ShipOnTime	string	`json:"ship_on_time"`
    MainCategoryId	int	`json:"main_category_id"`
    PositiveSellerRating	string	`json:"positive_seller_rating"`
}
func (g GetSellerMetricsByIdDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
