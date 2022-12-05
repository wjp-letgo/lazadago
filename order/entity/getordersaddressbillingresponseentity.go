package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOrdersAddressBillingResponseEntity struct {
	Address1  string `json:"address1"`
	Phone2    string `json:"phone2"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
	Address5  string `json:"address5"`
	PostCode  string `json:"post_code"`
	Address4  string `json:"address4"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	Address3  string `json:"address3"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
}

func (g GetOrdersAddressBillingResponseEntity) String() string {
	return lib.ObjectToString(g)
}
