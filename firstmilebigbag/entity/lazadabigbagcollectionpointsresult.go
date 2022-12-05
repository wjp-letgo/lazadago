package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCollectionPointsResult struct{
    Result	LazadaBigbagCollectionPointsResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]LazadaBigbagCollectionPointsDetailResponseEntity	`json:"detail"`
}
func (g LazadaBigbagCollectionPointsResult) String() string {
    return lib.ObjectToString(g)
}
type LazadaBigbagCollectionPointsDetailResponseEntity struct{}
