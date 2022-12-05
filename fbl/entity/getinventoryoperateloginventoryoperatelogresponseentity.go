package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetInventoryOperateLogInventoryOperateLogResponseEntity struct{
    RefOrderCode	[]GetInventoryOperateLogRefOrderCodeResponseEntity	`json:"ref_order_code"`
    WarehouseCode	string	`json:"warehouse_code"`
    WarehouseName	string	`json:"warehouse_name"`
    OrderType	string	`json:"order_type"`
    InventoryType	string	`json:"inventory_type"`
    ChangeQuantity	string	`json:"change_quantity"`
    ResultQuantity	string	`json:"result_quantity"`
    OperateTime	string	`json:"operate_time"`
}
func (g GetInventoryOperateLogInventoryOperateLogResponseEntity) String() string {
    return lib.ObjectToString(g)
}
