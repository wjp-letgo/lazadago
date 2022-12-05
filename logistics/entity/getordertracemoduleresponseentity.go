package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOrderTraceModuleResponseEntity struct {
	WarehouseDetailInfo   string                                             `json:"warehouse_detail_info"`
	OfcOrderId            int64                                              `json:"ofc_order_id"`
	PackageDetailInfoList []GetOrderTracePackageDetailInfoListResponseEntity `json:"package_detail_info_list"`
}

func (g GetOrderTraceModuleResponseEntity) String() string {
	return lib.ObjectToString(g)
}
