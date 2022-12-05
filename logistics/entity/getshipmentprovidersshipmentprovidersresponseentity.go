package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShipmentProvidersShipmentProvidersResponseEntity struct{
    IsDefault	int	`json:"is_default"`
    TrackingCodeExample	string	`json:"tracking_code_example"`
    EnabledDeliveryOptions	string	`json:"enabled_delivery_options"`
    Name	string	`json:"name"`
    Cod	int	`json:"cod"`
    TrackingCodeValidationRegex	string	`json:"tracking_code_validation_regex"`
    TrackingUrl	string	`json:"tracking_url"`
    ApiIntegration	int	`json:"api_integration"`
}
func (g GetShipmentProvidersShipmentProvidersResponseEntity) String() string {
    return lib.ObjectToString(g)
}
