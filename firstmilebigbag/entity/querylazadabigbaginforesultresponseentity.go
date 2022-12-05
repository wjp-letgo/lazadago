package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryLazadaBigbagInfoResultResponseEntity struct{
    Data	QueryLazadaBigbagInfoDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"errorCode"`
    ErrorMsg	string	`json:"errorMsg"`
}
func (g QueryLazadaBigbagInfoResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type QueryLazadaBigbagInfoDataResponseEntity struct{}
func (g QueryLazadaBigbagInfoDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}