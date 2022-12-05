package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrdersDataResponseEntity struct{
    CountTotal	int	`json:"countTotal"`
    Count	int	`json:"count"`
    Orders	[]GetOrdersOrdersResponseEntity	`json:"orders"`
}
func (g GetOrdersDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
