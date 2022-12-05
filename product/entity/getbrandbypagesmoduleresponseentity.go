package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBrandByPagesModuleResponseEntity struct{
    GlobalIdentifier	string	`json:"global_identifier"`
    NameEn	string	`json:"name_en"`
    BrandId	int	`json:"brand_id"`
    Name	string	`json:"name"`
}
func (g GetBrandByPagesModuleResponseEntity) String() string {
    return lib.ObjectToString(g)
}
