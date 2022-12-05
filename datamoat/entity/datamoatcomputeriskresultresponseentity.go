package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DataMoatComputeRiskResultResponseEntity struct{
    Msg	string	`json:"msg"`
    Success	bool	`json:"success"`
    Risk	string	`json:"risk"`
    RiskType	string	`json:"riskType"`
    RiskDescription	string	`json:"riskDescription"`
}
func (g DataMoatComputeRiskResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
