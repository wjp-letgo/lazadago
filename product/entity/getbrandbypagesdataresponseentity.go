package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBrandByPagesDataResponseEntity struct{
    StartRow	int	`json:"start_row"`
    PageIndex	int	`json:"page_index"`
    TotalPage	int	`json:"total_page"`
    Module	[]GetBrandByPagesModuleResponseEntity	`json:"module"`
    EnableTotal	bool	`json:"enable_total"`
    PageSize	int	`json:"page_size"`
    TotalRecord	int	`json:"total_record"`
}
func (g GetBrandByPagesDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
