package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToPackedByMarketplaceOrderItemsResponseEntity struct{
    OrderItemId	int64	`json:"order_item_id"`
    PurchaseOrderId	int64	`json:"purchase_order_id"`
    PurchaseOrderNumber	string	`json:"purchase_order_number"`
    PackageId	string	`json:"package_id"`
    ShipmentProvider	string	`json:"shipment_provider"`
    TrackingNumber	string	`json:"tracking_number"`
}
func (g SetStatusToPackedByMarketplaceOrderItemsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
