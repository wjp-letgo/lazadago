package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DataMoatComputeRiskResult struct{
    Result	DataMoatComputeRiskResultResponseEntity	`json:"result"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]DataMoatComputeRiskDetailResponseEntity	`json:"detail"`
}
func (g DataMoatComputeRiskResult) String() string {
    return lib.ObjectToString(g)
}
type DataMoatComputeRiskDetailResponseEntity struct{}
