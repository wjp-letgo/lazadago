package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFailureReasonsDataResponseEntity struct{
    ReasonId	int	`json:"reason_id"`
    Type	string	`json:"type"`
    Name	string	`json:"name"`
}
func (g GetFailureReasonsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
