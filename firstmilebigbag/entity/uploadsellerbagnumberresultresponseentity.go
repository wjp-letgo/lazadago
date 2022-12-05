package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadSellerBagNumberResultResponseEntity struct{
    Successs	bool	`json:"successs"`
    Data	UploadSellerBagNumberDataResponseEntity	`json:"data"`
    ResponseMessage	string	`json:"response_message"`
    ResponseCode	int	`json:"response_code"`
}
func (g UploadSellerBagNumberResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
