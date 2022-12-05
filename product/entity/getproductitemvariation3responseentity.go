package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductItemVariation3ResponseEntity struct{
    Name	string	`json:"name"`
    HasImage	bool	`json:"has_image"`
    Customize	bool	`json:"customize"`
    Options	[]string	`json:"options"`
}
func (g GetProductItemVariation3ResponseEntity) String() string {
    return lib.ObjectToString(g)
}
