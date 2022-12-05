package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagUpdateResultResponseEntity struct{
    Data	LazadaBigbagUpdateDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"errorCode"`
    ErroMsg	string	`json:"erroMsg"`
}
func (g LazadaBigbagUpdateResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagUpdateDataResponseEntity struct{}
func (g LazadaBigbagUpdateDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}