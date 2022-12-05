package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderTracePackageDetailInfoListResponseEntity struct{
    OrderLineInfoList	string	`json:"order_line_info_list"`
    TrackingNumber	string	`json:"tracking_number"`
    OfcPackageId	string	`json:"ofc_package_id"`
    LogisticDetailInfoList	[]GetOrderTraceLogisticDetailInfoListResponseEntity	`json:"logistic_detail_info_list"`
}
func (g GetOrderTracePackageDetailInfoListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
