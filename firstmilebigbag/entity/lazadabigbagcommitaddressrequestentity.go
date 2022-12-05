package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagCommitAddressRequestEntity struct{
    Country	string	`json:"country"`
    ZipCode	string	`json:"zipCode"`
    City	string	`json:"city"`
    Province	string	`json:"province"`
    Street	string	`json:"street"`
    District	string	`json:"district"`
    DetailAddress	string	`json:"detailAddress"`
}
func (g LazadaBigbagCommitAddressRequestEntity) String() string {
    return lib.ObjectToString(g)
}
