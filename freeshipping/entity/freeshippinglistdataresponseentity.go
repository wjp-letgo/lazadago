package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type FreeShippingListDataResponseEntity struct {
	Total    int                                      `json:"total"`
	Current  int                                      `json:"current"`
	DataList []FreeShippingListDataListResponseEntity `json:"data_list"`
	PageSize int                                      `json:"page_size"`
}

func (g FreeShippingListDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
