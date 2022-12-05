package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMultiWarehouseBySellerResultResponseEntity struct {
	Success   bool                                             `json:"success"`
	Module    GetMultiWarehouseBySellerModuleResponseEntity    `json:"module"`
	ErrorCode GetMultiWarehouseBySellerErrorCodeResponseEntity `json:"error_code"`
}

func (g GetMultiWarehouseBySellerResultResponseEntity) String() string {
	return lib.ObjectToString(g)
}

type GetMultiWarehouseBySellerModuleResponseEntity struct{}
