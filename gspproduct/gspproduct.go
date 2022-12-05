package gspproduct

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	gspproductentity "github.com/wjpxxx/lazadago/gspproduct/entity"
)
//GspProduct
type GspProduct struct{
	Config *lazadaConfig.Config
}

//CreateGlobalProduct
//@Title Use this API to create a single new global product to multiple Lazada sites.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=18&path=/product/global/create
func (s *GspProduct) CreateGlobalProduct (payload string) gspproductentity.CreateGlobalProductResult {
    method := "/product/global/create"
    params := lib.InRow{
      "payload":payload,
    }
    result := gspproductentity.CreateGlobalProductResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetGlobalProductStatus
//@Title Use this API to query the status of the specified global product. It takes several minutes for the global product to be created on each site.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=18&path=/product/global/status/get
func (s *GspProduct) GetGlobalProductStatus (p gspproductentity.GetGlobalProductStatusParamsRequestEntity) gspproductentity.GetGlobalProductStatusResult {
    method := "/product/global/status/get"
    params := lib.InRow{
      "params":p,
    }
    result := gspproductentity.GetGlobalProductStatusResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetUnfilledAttribute
//@Title get the product which hava attribute not filled
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=18&path=/product/global/unfilled/attribute/get
func (s *GspProduct) GetUnfilledAttribute (offset int,limit int,attributeTag string) gspproductentity.GetUnfilledAttributeResult {
    method := "/product/global/unfilled/attribute/get"
    params := lib.InRow{
      "offset":offset,
      "limit":limit,
      "attributeTag":attributeTag,
    }
    result := gspproductentity.GetUnfilledAttributeResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UpdateGlobalProductAttribute
//@Title update global product attribute
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=18&path=/product/global/attribute/update
func (s *GspProduct) UpdateGlobalProductAttribute (payload string) gspproductentity.UpdateGlobalProductAttributeResult {
    method := "/product/global/attribute/update"
    params := lib.InRow{
      "payload":payload,
    }
    result := gspproductentity.UpdateGlobalProductAttributeResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
