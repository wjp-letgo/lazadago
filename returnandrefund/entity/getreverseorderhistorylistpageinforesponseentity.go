package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReverseOrderHistoryListPageInfoResponseEntity struct{
    PageSize	int	`json:"page_size"`
    CurrentPageNumber	int	`json:"current_page_number"`
    Total	int	`json:"total"`
}
func (g GetReverseOrderHistoryListPageInfoResponseEntity) String() string {
    return lib.ObjectToString(g)
}
