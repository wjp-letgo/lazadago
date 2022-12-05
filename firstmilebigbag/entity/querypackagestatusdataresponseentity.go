package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryPackageStatusDataResponseEntity struct{
    RtsDate	string	`json:"rts_date"`
    OrderCreateDate	string	`json:"order_create_date"`
    IsPickedUp	bool	`json:"is_picked_up"`
    OrderNumber	int64	`json:"order_number"`
    Weight	string	`json:"weight"`
    IsReceived	bool	`json:"is_received"`
    GlobalCollection	string	`json:"global_collection"`
    ReceiveDate	string	`json:"receive_date"`
    IsScanin	bool	`json:"is_scanin"`
    PickupDate	string	`json:"pickup_date"`
    TrackingNumber	string	`json:"tracking_number"`
    LzdBagNumber	string	`json:"lzd_bag_number"`
}
func (g QueryPackageStatusDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
