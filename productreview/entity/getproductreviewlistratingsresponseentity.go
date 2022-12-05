package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductReviewListRatingsResponseEntity struct{
    ProductRating	int	`json:"product_rating"`
    SellerRating	int	`json:"seller_rating"`
    LogisticsRating	int	`json:"logistics_rating"`
}
func (g GetProductReviewListRatingsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
