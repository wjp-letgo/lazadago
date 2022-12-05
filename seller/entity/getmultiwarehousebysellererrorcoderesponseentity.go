package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMultiWarehouseBySellerErrorCodeResponseEntity struct{
    ErrorCodeParams	[]GetMultiWarehouseBySellerErrorCodeParamsResponseEntity	`json:"error_code_params"`
    DisplayMessage	string	`json:"display_message"`
}
func (g GetMultiWarehouseBySellerErrorCodeResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type GetMultiWarehouseBySellerErrorCodeParamsResponseEntity struct{}