package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCancelResultResponseEntity struct{
    Data	LazadaBigbagCancelDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"error_code"`
    ErrorMsg	string	`json:"error_msg"`
}
func (g LazadaBigbagCancelResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagCancelDataResponseEntity struct{}
func (g LazadaBigbagCancelDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}