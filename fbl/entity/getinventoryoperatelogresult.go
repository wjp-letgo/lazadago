package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInventoryOperateLogResult struct{
    InventoryOperateLog	[]GetInventoryOperateLogInventoryOperateLogResponseEntity	`json:"inventory_operate_log"`
    Success	string	`json:"success"`
    ErrMessage	string	`json:"errMessage"`
    ErrCode	string	`json:"errCode"`
    Page	int	`json:"page"`
    PerPage	int	`json:"per_page"`
    TotalCount	int	`json:"total_count"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetInventoryOperateLogDetailResponseEntity	`json:"detail"`
}
func (g GetInventoryOperateLogResult) String() string {
    return lib.ObjectToString(g)
}
type GetInventoryOperateLogDetailResponseEntity struct{}
