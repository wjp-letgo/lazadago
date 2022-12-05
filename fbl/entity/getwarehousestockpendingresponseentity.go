package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockPendingResponseEntity struct{
    Reserved	int	`json:"reserved"`
    Available	int	`json:"available"`
}
func (g GetWarehouseStockPendingResponseEntity) String() string {
    return lib.ObjectToString(g)
}
