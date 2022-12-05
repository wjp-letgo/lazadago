package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOutboundOrderListResultResponseEntity struct{
    PerPage	int	`json:"per_page"`
    Data	[]GetOutboundOrderListDataResponseEntity	`json:"data"`
    Page	int	`json:"page"`
    TotalCount	int	`json:"total_count"`
}
func (g GetOutboundOrderListResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
