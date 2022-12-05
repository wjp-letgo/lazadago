package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboDataResponseEntity struct{
    PageSize	int	`json:"page_size"`
    Total	int	`json:"total"`
    Current	int	`json:"current"`
    DataList	[]ListFlexiComboDataListResponseEntity	`json:"data_list"`
}
func (g ListFlexiComboDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
