package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type LazadaBigbagCommitCollectionInfoRequestEntity struct {
	PickUpCode string `json:"pickUpCode"`
}

func (g LazadaBigbagCommitCollectionInfoRequestEntity) String() string {
	return lib.ObjectToString(g)
}
