package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateProductVariation4ResponseEntity struct {
	Name      string   `json:"name"`
	HasImage  bool     `json:"has_image"`
	Customize bool     `json:"customize"`
	Options   []string `json:"options"`
}

func (g UpdateProductVariation4ResponseEntity) String() string {
	return lib.ObjectToString(g)
}
