package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrdersForSellerResultResponseEntity struct{
    Total	int	`json:"total"`
    Items	[]GetReverseOrdersForSellerItemsResponseEntity	`json:"items"`
    PageNo	int	`json:"page_no"`
    Success	bool	`json:"success"`
    PageSize	int	`json:"page_size"`
}
func (g GetReverseOrdersForSellerResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
