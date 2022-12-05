package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductReviewListDataResponseEntity struct{
    Data	[]GetProductReviewListDataResponseEntity	`json:"data"`
    Current	int	`json:"current"`
    PageSize	int	`json:"page_size"`
    Total	int	`json:"total"`
}
func (g GetProductReviewListDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
