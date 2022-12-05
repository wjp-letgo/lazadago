package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInventoryChangedSKUResult struct{
    PerPage	int	`json:"per_page"`
    Page	int	`json:"page"`
    TotalCount	int	`json:"total_count"`
    SkuList	[]GetInventoryChangedSKUSkuListResponseEntity	`json:"sku_list"`
    Success	string	`json:"success"`
    ErrMessage	string	`json:"errMessage"`
    ErrCode	string	`json:"errCode"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetInventoryChangedSKUDetailResponseEntity	`json:"detail"`
}
func (g GetInventoryChangedSKUResult) String() string {
    return lib.ObjectToString(g)
}
type GetInventoryChangedSKUDetailResponseEntity struct{}
