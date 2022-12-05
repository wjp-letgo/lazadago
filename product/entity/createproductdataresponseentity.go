package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type CreateProductDataResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    SkuList	[]CreateProductSkuListResponseEntity	`json:"sku_list"`
    Variation1	CreateProductVariation1ResponseEntity	`json:"Variation1"`
    Variation2	CreateProductVariation2ResponseEntity	`json:"Variation2"`
    Variation3	CreateProductVariation3ResponseEntity	`json:"Variation3"`
    Variation4	CreateProductVariation4ResponseEntity	`json:"Variation4"`
}
func (g CreateProductDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
