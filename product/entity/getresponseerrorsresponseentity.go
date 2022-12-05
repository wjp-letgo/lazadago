package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetResponseErrorsResponseEntity struct{
    Field	string	`json:"field"`
    Msg	string	`json:"msg"`
}
func (g GetResponseErrorsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
