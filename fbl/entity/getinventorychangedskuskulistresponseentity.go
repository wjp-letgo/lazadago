package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInventoryChangedSKUSkuListResponseEntity struct{
    FulfillmentSkuId	int64	`json:"fulfillment_sku_id"`
    OperateLogCount	int	`json:"operate_log_count"`
}
func (g GetInventoryChangedSKUSkuListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
