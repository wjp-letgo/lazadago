package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadLazadaBagNumberResultResponseEntity struct{
    ResponseCode	int	`json:"response_code"`
    Successs	UploadLazadaBagNumberSuccesssResponseEntity	`json:"successs"`
    Data	UploadLazadaBagNumberDataResponseEntity	`json:"data"`
    ResponseMessage	string	`json:"response_message"`
}
func (g UploadLazadaBagNumberResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type UploadLazadaBagNumberSuccesssResponseEntity struct{}
func (g UploadLazadaBagNumberSuccesssResponseEntity) String() string {
    return lib.ObjectToString(g)
}
