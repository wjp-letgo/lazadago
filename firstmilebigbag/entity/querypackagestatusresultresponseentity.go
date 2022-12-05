package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type QueryPackageStatusResultResponseEntity struct{
    Successs	bool	`json:"successs"`
    Data	[]QueryPackageStatusDataResponseEntity	`json:"data"`
    ResponseMessage	string	`json:"response_message"`
    ResponseCode	int	`json:"response_code"`
}
func (g QueryPackageStatusResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
