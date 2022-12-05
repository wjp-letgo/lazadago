package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetResponseDataResponseEntity struct{
    Images	[]GetResponseImagesResponseEntity	`json:"images"`
    Errors	[]GetResponseErrorsResponseEntity	`json:"errors"`
}
func (g GetResponseDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
