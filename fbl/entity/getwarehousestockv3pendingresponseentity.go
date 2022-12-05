package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWarehouseStockV3PendingResponseEntity struct{
    Available	int	`json:"available"`
    Reserved	int	`json:"reserved"`
}
func (g GetWarehouseStockV3PendingResponseEntity) String() string {
    return lib.ObjectToString(g)
}
