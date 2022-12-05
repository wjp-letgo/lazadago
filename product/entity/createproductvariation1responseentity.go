package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type CreateProductVariation1ResponseEntity struct {
	Name      string   `json:"name"`
	HasImage  bool     `json:"has_image"`
	Customize bool     `json:"customize"`
	Options   []string `json:"options"`
}

func (g CreateProductVariation1ResponseEntity) String() string {
	return lib.ObjectToString(g)
}
