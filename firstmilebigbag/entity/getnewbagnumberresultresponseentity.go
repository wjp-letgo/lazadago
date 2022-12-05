package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetNewBagNumberResultResponseEntity struct{
    Successs	bool	`json:"successs"`
    Data	GetNewBagNumberDataResponseEntity	`json:"data"`
    ResponseMessage	string	`json:"response_message"`
    ResponseCode	int	`json:"response_code"`
}
func (g GetNewBagNumberResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
