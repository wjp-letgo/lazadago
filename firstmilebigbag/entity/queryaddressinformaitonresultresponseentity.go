package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryAddressInformaitonResultResponseEntity struct{
    Data	QueryAddressInformaitonDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"errorCode"`
    ErrorMsg	string	`json:"errorMsg"`
}
func (g QueryAddressInformaitonResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type QueryAddressInformaitonDataResponseEntity struct{}
func (g QueryAddressInformaitonDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
