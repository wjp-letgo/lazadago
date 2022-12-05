package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetCategoryAttributesDataResponseEntity struct{
    Advanced	GetCategoryAttributesAdvancedResponseEntity	`json:"advanced"`
    Label	string	`json:"label"`
    Name	string	`json:"name"`
    IsMandatory	int	`json:"is_mandatory"`
    AttributeType	string	`json:"attribute_type"`
    InputType	string	`json:"input_type"`
    Options	[]GetCategoryAttributesOptionsResponseEntity	`json:"options"`
    IsSaleProp	int	`json:"is_sale_prop"`
}
func (g GetCategoryAttributesDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type GetCategoryAttributesAdvancedResponseEntity struct{
    IsKeyProp int `json:"is_key_prop"`
}
func (g GetCategoryAttributesAdvancedResponseEntity) String() string {
    return lib.ObjectToString(g)
}