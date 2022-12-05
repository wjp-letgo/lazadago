package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnfilledAttributeAdvancedResponseEntity struct{
    IsKeyProp	int	`json:"is_key_prop"`
}
func (g GetUnfilledAttributeAdvancedResponseEntity) String() string {
    return lib.ObjectToString(g)
}
