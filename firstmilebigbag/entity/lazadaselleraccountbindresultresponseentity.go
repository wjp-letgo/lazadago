package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaSellerAccountBindResultResponseEntity struct{
    Data	LazadaSellerAccountBindDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"errorCode"`
    ErrorMsg	string	`json:"errorMsg"`
}
func (g LazadaSellerAccountBindResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type LazadaSellerAccountBindDataResponseEntity struct{}
func (g LazadaSellerAccountBindDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}