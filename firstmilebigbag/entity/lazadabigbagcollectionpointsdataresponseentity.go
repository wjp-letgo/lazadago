package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCollectionPointsDataResponseEntity struct{
    CurrentPageIndex	int	`json:"currentPageIndex"`
    PageTotalNum	int	`json:"pageTotalNum"`
    PageSize	int	`json:"pageSize"`
    TotalCount	int	`json:"totalCount"`
    ItemList	[]LazadaBigbagCollectionPointsItemListResponseEntity	`json:"itemList"`
}
func (g LazadaBigbagCollectionPointsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagCollectionPointsItemListResponseEntity struct{}
func (g LazadaBigbagCollectionPointsItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}