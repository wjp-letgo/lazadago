package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateProductVariationResponseEntity struct{
    Variation1	UpdateProductVariation1ResponseEntity	`json:"Variation1"`
    Variation2	UpdateProductVariation2ResponseEntity	`json:"Variation2"`
    Variation3	UpdateProductVariation3ResponseEntity	`json:"Variation3"`
    Variation4	UpdateProductVariation4ResponseEntity	`json:"Variation4"`
}
func (g UpdateProductVariationResponseEntity) String() string {
    return lib.ObjectToString(g)
}
