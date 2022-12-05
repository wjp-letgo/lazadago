package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RetailFulfilmentUpdateFeaturesRequestEntity struct{
    SKUSELLERNICODE	string	`json:"SKUSELLERNICODE"`
    RMCATEGORY1	string	`json:"RM_CATEGORY1"`
    RMCATEGORY2	string	`json:"RM_CATEGORY2"`
    RMCATEGORY3	string	`json:"RM_CATEGORY3"`
    RMTRANSPCS	string	`json:"RM_TRANSPCS"`
    RMSELLINGUNIT	string	`json:"RM_SELLINGUNIT"`
    RMSELLINGPCS	string	`json:"RM_SELLINGPCS"`
    RMMINDELIACPDAY	string	`json:"RM_MINDELIACPDAY"`
    RMWHLOCATIONZONE	string	`json:"RM_WHLOCATIONZONE"`
    RMSTATUSREASON	string	`json:"RM_STATUSREASON"`
    RMSHELFLIFEGUARANTEE	string	`json:"RM_SHELFLIFEGUARANTEE"`
    RMTRANSUNIT	string	`json:"RM_TRANSUNIT"`
    RMTIERSPERPALLET	string	`json:"RM_TIERSPERPALLET"`
    RMUNITSPERTIERS	string	`json:"RM_UNITSPERTIERS"`
    RMSTATUS	string	`json:"RM_STATUS"`
    RMSUPPLIERNAME	string	`json:"RM_SUPPLIERNAME"`
    RMSUPPLIERPURCHASEPOINT	string	`json:"RM_SUPPLIERPURCHASEPOINT"`
    RMITEMMOQ	string	`json:"RM_ITEMMOQ"`
    RMSUPPLYHOLDREASON	string	`json:"RM_SUPPLYHOLDREASON"`
    RMSUPPLYHOLDREASONENDDATE	string	`json:"RM_SUPPLYHOLDREASONENDDATE"`
    RMFIRSTORDERQUANTITY	string	`json:"RM_FIRSTORDERQUANTITY"`
    RMSUPPLIERPRODUCTCODE	string	`json:"RM_SUPPLIERPRODUCTCODE"`
    RMMEASUREVALUE	string	`json:"RM_MEASUREVALUE"`
    RMMEASUREMETRIC	string	`json:"RM_MEASUREMETRIC"`
    RMCOMMERCIAL	string	`json:"RM_COMMERCIAL"`
    VAT	string	`json:"VAT"`
}
func (g RetailFulfilmentUpdateFeaturesRequestEntity) String() string {
    return lib.ObjectToString(g)
}
