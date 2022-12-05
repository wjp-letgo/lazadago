package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeAttributesResponseEntity struct{
    Advanced	GetUnfilledAttributeAdvancedResponseEntity	`json:"advanced"`
    InputType	string	`json:"input_type"`
    Options	[]string	`json:"options"`
    Name	string	`json:"name"`
    IsMandatory	int	`json:"is_mandatory"`
    AttributeType	string	`json:"attribute_type"`
    Label	string	`json:"label"`
}
func (g GetUnfilledAttributeAttributesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
