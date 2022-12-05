package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategoryAttributesOptionsResponseEntity struct{
    Name	string	`json:"name"`
}
func (g GetCategoryAttributesOptionsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
