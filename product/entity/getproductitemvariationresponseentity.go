package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductItemVariationResponseEntity struct{
    Variation1	GetProductItemVariation1ResponseEntity	`json:"variation1"`
    Variation2	GetProductItemVariation2ResponseEntity	`json:"variation2"`
    Variation3	GetProductItemVariation3ResponseEntity	`json:"variation3"`
    Variation4	GetProductItemVariation4ResponseEntity	`json:"variation4"`
}
func (g GetProductItemVariationResponseEntity) String() string {
    return lib.ObjectToString(g)
}
