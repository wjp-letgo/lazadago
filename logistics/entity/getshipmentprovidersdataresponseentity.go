package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetShipmentProvidersDataResponseEntity struct {
	ShipmentProviders []GetShipmentProvidersShipmentProvidersResponseEntity `json:"shipment_providers"`
}

func (g GetShipmentProvidersDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
