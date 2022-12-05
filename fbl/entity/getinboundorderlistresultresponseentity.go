package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInboundOrderListResultResponseEntity struct{
    PerPage	int	`json:"per_page"`
    Data	[]GetInboundOrderListDataResponseEntity	`json:"data"`
    Page	int	`json:"page"`
    TotalCount	int	`json:"total_count"`
}
func (g GetInboundOrderListResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
