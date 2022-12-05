package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCommitResultResponseEntity struct{
    Data	LazadaBigbagCommitDataResponseEntity	`json:"data"`
    Success	bool	`json:"success"`
    ErrorCode	string	`json:"errorCode"`
    ErrorMsg	string	`json:"errorMsg"`
}
func (g LazadaBigbagCommitResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
