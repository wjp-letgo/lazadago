package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetReverseOrderHistoryListDataResponseEntity struct {
	List     []GetReverseOrderHistoryListListResponseEntity   `json:"list"`
	PageInfo GetReverseOrderHistoryListPageInfoResponseEntity `json:"page_info"`
}

func (g GetReverseOrderHistoryListDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
