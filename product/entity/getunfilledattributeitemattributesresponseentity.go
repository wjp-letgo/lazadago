package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeItemAttributesResponseEntity struct{
    Advanced	GetUnfilledAttributeItemAdvancedResponseEntity	`json:"advanced"`
    Name	string	`json:"name"`
    InputType	string	`json:"input_type"`
    Options	[]GetUnfilledAttributeItemOptionsResponseEntity	`json:"options"`
    IsMandatory	int	`json:"is_mandatory"`
    AttributeType	string	`json:"attribute_type"`
    Label	string	`json:"label"`
}
func (g GetUnfilledAttributeItemAttributesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
//GetUnfilledAttributeItemAdvancedResponseEntity
type GetUnfilledAttributeItemAdvancedResponseEntity struct{
    IsKeyProp int `json:"is_key_prop"`
}
func (g GetUnfilledAttributeItemAdvancedResponseEntity) String() string {
    return lib.ObjectToString(g)
}