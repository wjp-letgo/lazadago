package freeshipping

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	freeshippingentity "github.com/wjpxxx/lazadago/freeshipping/entity"
)
//FreeShipping
type FreeShipping struct{
	Config *lazadaConfig.Config
}

//FreeShippingActivate
//@Title activate free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/activate
func (s *FreeShipping) FreeShippingActivate (id int64) freeshippingentity.FreeShippingActivateResult {
    method := "/promotion/freeshipping/activate"
    params := lib.InRow{
      "id":id,
    }
    result := freeshippingentity.FreeShippingActivateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingAddSelectedProductSKU
//@Title add sku for free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/product/sku/add
func (s *FreeShipping) FreeShippingAddSelectedProductSKU (id int64,skuIds []int64) freeshippingentity.FreeShippingAddSelectedProductSKUResult {
    method := "/promotion/freeshipping/product/sku/add"
    params := lib.InRow{
      "id":id,
      "sku_ids":skuIds,
    }
    result := freeshippingentity.FreeShippingAddSelectedProductSKUResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingCreate
//@Title create a new free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/create
func (s *FreeShipping) FreeShippingCreate (budgetType string,templateType string,apply string,periodEndTime int,templateCode string,categoryName string,budgetValue string,promotionName string,periodType string,regionType string,periodStartTime int,platformChannel string,campaignTag string,regionValue []string,deliveryOption string,tiers []freeshippingentity.FreeShippingCreateTiersRequestEntity,discountType string,dealCriteria string) freeshippingentity.FreeShippingCreateResult {
    method := "/promotion/freeshipping/create"
    params := lib.InRow{
      "budget_type":budgetType,
      "apply":apply,
      "period_end_time":periodEndTime,
      "promotion_name":promotionName,
      "period_type":periodType,
      "region_type":regionType,
      "period_start_time":periodStartTime,
      "delivery_option":deliveryOption,
      "tiers":tiers,
      "discount_type":discountType,
      "deal_criteria":dealCriteria,
    }
    if templateType!=""{
        params["template_type"]=templateType
    }
    if templateCode!=""{
        params["template_code"]=templateCode
    }
    if categoryName!=""{
        params["category_name"]=categoryName
    }
    if budgetValue!=""{
        params["budget_value"]=budgetValue
    }
    if platformChannel!=""{
        params["platform_channel"]=platformChannel
    }
    if campaignTag!=""{
        params["campaign_tag"]=campaignTag
    }
    result := freeshippingentity.FreeShippingCreateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingDeactivate
//@Title deactivate free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/deactivate
func (s *FreeShipping) FreeShippingDeactivate (id int64) freeshippingentity.FreeShippingDeactivateResult {
    method := "/promotion/freeshipping/deactivate"
    params := lib.InRow{
      "id":id,
    }
    result := freeshippingentity.FreeShippingDeactivateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingDeleteSelectedProductSKU
//@Title delete sku for free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/product/sku/remove
func (s *FreeShipping) FreeShippingDeleteSelectedProductSKU (id int64,skuIds []int64) freeshippingentity.FreeShippingDeleteSelectedProductSKUResult {
    method := "/promotion/freeshipping/product/sku/remove"
    params := lib.InRow{
      "id":id,
      "sku_ids":skuIds,
    }
    result := freeshippingentity.FreeShippingDeleteSelectedProductSKUResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingDeliveryOptionsQuery
//@Title query free shipping promotion delivery options
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/deliveryoptions/get
func (s *FreeShipping) FreeShippingDeliveryOptionsQuery () freeshippingentity.FreeShippingDeliveryOptionsQueryResult {
    method := "/promotion/freeshipping/deliveryoptions/get"
    params := lib.InRow{}
    result := freeshippingentity.FreeShippingDeliveryOptionsQueryResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingGet
//@Title get free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/get
func (s *FreeShipping) FreeShippingGet (id int64) freeshippingentity.FreeShippingGetResult {
    method := "/promotion/freeshipping/get"
    params := lib.InRow{
      "id":id,
    }
    result := freeshippingentity.FreeShippingGetResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingList
//@Title query free shipping promotion list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshippings/get
func (s *FreeShipping) FreeShippingList (curPage int,name string,pageSize int,status string) freeshippingentity.FreeShippingListResult {
    method := "/promotion/freeshippings/get"
    params := lib.InRow{
    }
    if curPage>-1{
        params["curPage"]=curPage
    }
    if name!=""{
        params["name"]=name
    }
    if pageSize>-1{
        params["pageSize"]=pageSize
    }
    if status!=""{
        params["status"]=status
    }
    result := freeshippingentity.FreeShippingListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingRegionsQuery
//@Title query free shipping promotion regions
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/regions/get
func (s *FreeShipping) FreeShippingRegionsQuery () freeshippingentity.FreeShippingRegionsQueryResult {
    method := "/promotion/freeshipping/regions/get"
    params := lib.InRow{}
    result := freeshippingentity.FreeShippingRegionsQueryResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingSelectedProductList
//@Title query free shipping promotion selected product list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/products/get
func (s *FreeShipping) FreeShippingSelectedProductList (curPage int,pageSize int,id int64) freeshippingentity.FreeShippingSelectedProductListResult {
    method := "/promotion/freeshipping/products/get"
    params := lib.InRow{
      "id":id,
    }
    if curPage>-1{
        params["curPage"]=curPage
    }
    if pageSize>-1{
        params["pageSize"]=pageSize
    }
    result := freeshippingentity.FreeShippingSelectedProductListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//FreeShippingUpdate
//@Title update free shipping promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=31&path=/promotion/freeshipping/update
func (s *FreeShipping) FreeShippingUpdate (budgetType string,templateType string,apply string,periodEndTime int,templateCode string,categoryName string,budgetValue string,promotionName string,periodType string,regionType string,periodStartTime int,platformChannel string,campaignTag string,regionValue []string,id int64,deliveryOption string,tiers []freeshippingentity.FreeShippingUpdateTiersRequestEntity,discountType string,dealCriteria string) freeshippingentity.FreeShippingUpdateResult {
    method := "/promotion/freeshipping/update"
    params := lib.InRow{
      "budget_type":budgetType,
      "template_type":templateType,
      "apply":apply,
      "period_end_time":periodEndTime,
      "promotion_name":promotionName,
      "period_type":periodType,
      "region_type":regionType,
      "period_start_time":periodStartTime,
      "id":id,
      "delivery_option":deliveryOption,
      "tiers":tiers,
      "discount_type":discountType,
      "deal_criteria":dealCriteria,
    }
    if templateCode!=""{
        params["template_code"]=templateCode
    }
    if categoryName!=""{
        params["category_name"]=categoryName
    }
    if budgetValue!=""{
        params["budget_value"]=budgetValue
    }
    if platformChannel!=""{
        params["platform_channel"]=platformChannel
    }
    if campaignTag!=""{
        params["campaign_tag"]=campaignTag
    }
    result := freeshippingentity.FreeShippingUpdateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
