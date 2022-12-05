package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderTraceResultResponseEntity struct{
    ErrorCode	GetOrderTraceErrorCodeResponseEntity	`json:"error_code"`
    Repeated	bool	`json:"repeated"`
    Retry	bool	`json:"retry"`
    NotSuccess	bool	`json:"not_success"`
    Success	bool	`json:"success"`
    Module	[]GetOrderTraceModuleResponseEntity	`json:"module"`
}
func (g GetOrderTraceResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
