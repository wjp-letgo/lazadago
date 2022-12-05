package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCommitDataResponseEntity struct{
    HandoverOrderId	int	`json:"handoverOrderId"`
    HandoverContentId	int	`json:"handoverContentId"`
    HandoverContentCode	string	`json:"handoverContentCode"`
}
func (g LazadaBigbagCommitDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
