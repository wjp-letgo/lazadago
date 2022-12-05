package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetSellerItemLimitDataResponseEntity struct {
	OnlineItemCount int `json:"onlineItemCount"`
	ItemLimit       int `json:"itemLimit"`
	PayItemCnt      int `json:"payItemCnt"`
	PayByrCnt       int `json:"payByrCnt"`
}

func (g GetSellerItemLimitDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
