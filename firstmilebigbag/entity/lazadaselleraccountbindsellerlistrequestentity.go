package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaSellerAccountBindSellerListRequestEntity struct {
	Country    string `json:"country"`
	SellerId   string `json:"sellerId"`
	ShortCode  string `json:"shortCode"`
	SellerName string `json:"sellerName"`
}

func (g LazadaSellerAccountBindSellerListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
