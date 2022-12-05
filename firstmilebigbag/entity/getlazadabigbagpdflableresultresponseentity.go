package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetLazadaBigbagPDFLableResultResponseEntity struct{
    Data	GetLazadaBigbagPDFLableDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"errorCode"`
    ErrorMsg	string	`json:"errorMsg"`
}
func (g GetLazadaBigbagPDFLableResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type GetLazadaBigbagPDFLableDataResponseEntity struct{}
func (g GetLazadaBigbagPDFLableDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}