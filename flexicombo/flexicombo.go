package flexicombo
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	flexicomboentity "github.com/wjpxxx/lazadago/flexicombo/entity"
)
//FlexiCombo
type FlexiCombo struct{
	Config *lazadaConfig.Config
}

//ActivateFlexiCombo
//@Title activate flexi combo
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/activate
func (s *FlexiCombo) ActivateFlexiCombo (id int64) flexicomboentity.ActivateFlexiComboResult {
    method := "/promotion/flexicombo/activate"
    params := lib.InRow{
      "id":id,
    }
    result := flexicomboentity.ActivateFlexiComboResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//AddFlexiComboProducts
//@Title add flexi combo products
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/products/add
func (s *FlexiCombo) AddFlexiComboProducts (id int64,skuIds []int64) flexicomboentity.AddFlexiComboProductsResult {
    method := "/promotion/flexicombo/products/add"
    params := lib.InRow{
      "id":id,
      "sku_ids":skuIds,
    }
    result := flexicomboentity.AddFlexiComboProductsResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//CreateFlexiCombo
//@Title create a  new promotion flexi combo
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/create
func (s *FlexiCombo) CreateFlexiCombo (apply string,sampleSkus []flexicomboentity.CreateFlexiComboSampleSkusRequestEntity,criteriaType string,criteriaValue []string,orderNumbers int,name string,platformChannel string,giftSkus []flexicomboentity.CreateFlexiComboGiftSkusRequestEntity,startTime int,discountType string,endTime int,discountValue []string,stackable string) flexicomboentity.CreateFlexiComboResult {
    method := "/promotion/flexicombo/create"
    params := lib.InRow{
      "apply":apply,
      "criteria_type":criteriaType,
      "criteria_value":criteriaValue,
      "order_numbers":orderNumbers,
      "name":name,
      "start_time":startTime,
      "discount_type":discountType,
      "end_time":endTime,
      "discount_value":discountValue,
    }
    if platformChannel!=""{
        params["platform_channel"]=platformChannel
    }
    if stackable!=""{
        params["stackable"]=stackable
    }
    result := flexicomboentity.CreateFlexiComboResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//DeactivateFlexiCombo
//@Title deactivate flexi combo
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/deactivate
func (s *FlexiCombo) DeactivateFlexiCombo (id int64) flexicomboentity.DeactivateFlexiComboResult {
    method := "/promotion/flexicombo/deactivate"
    params := lib.InRow{
      "id":id,
    }
    result := flexicomboentity.DeactivateFlexiComboResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//DeleteFlexiComboProducts
//@Title delete flexi combo products
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/products/delete
func (s *FlexiCombo) DeleteFlexiComboProducts (id int64,skuIds []int64) flexicomboentity.DeleteFlexiComboProductsResult {
    method := "/promotion/flexicombo/products/delete"
    params := lib.InRow{
      "id":id,
      "sku_ids":skuIds,
    }
    result := flexicomboentity.DeleteFlexiComboProductsResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetFlexiComboDetails
//@Title get promotion flexi combo detail by id
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/details
func (s *FlexiCombo) GetFlexiComboDetails (id int64) flexicomboentity.GetFlexiComboDetailsResult {
    method := "/promotion/flexicombo/details"
    params := lib.InRow{
      "id":id,
    }
    result := flexicomboentity.GetFlexiComboDetailsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//ListFlexiCombo
//@Title list flexi combo
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/list
func (s *FlexiCombo) ListFlexiCombo (curPage int,name string,pageSize int,status string) flexicomboentity.ListFlexiComboResult {
    method := "/promotion/flexicombo/list"
    params := lib.InRow{
      "cur_page":curPage,
      "page_size":pageSize,
    }
    if name!=""{
        params["name"]=name
    }
    if status!=""{
        params["status"]=status
    }
    result := flexicomboentity.ListFlexiComboResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//ListFlexiComboProducts
//@Title list flexi combo products
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/products/list
func (s *FlexiCombo) ListFlexiComboProducts (curPage int,pageSize int,id int64) flexicomboentity.ListFlexiComboProductsResult {
    method := "/promotion/flexicombo/products/list"
    params := lib.InRow{
      "cur_page":curPage,
      "page_size":pageSize,
      "id":id,
    }
    result := flexicomboentity.ListFlexiComboProductsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UpdateFlexiCombo
//@Title update flexi combo
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=25&path=/promotion/flexicombo/update
func (s *FlexiCombo) UpdateFlexiCombo (apply string,sampleSkus []flexicomboentity.UpdateFlexiComboSampleSkusRequestEntity,criteriaType string,criteriaValue []string,orderNumbers int,name string,platformChannel string,giftSkus []flexicomboentity.UpdateFlexiComboGiftSkusRequestEntity,startTime int,discountType string,id int64,endTime int,discountValue []string,stackable string) flexicomboentity.UpdateFlexiComboResult {
    method := "/promotion/flexicombo/update"
    params := lib.InRow{
      "apply":apply,
      "criteria_type":criteriaType,
      "criteria_value":criteriaValue,
      "order_numbers":orderNumbers,
      "name":name,
      "start_time":startTime,
      "discount_type":discountType,
      "id":id,
      "end_time":endTime,
      "discount_value":discountValue,
    }
    if platformChannel!=""{
        params["platform_channel"]=platformChannel
    }
    if stackable!=""{
        params["stackable"]=stackable
    }
    result := flexicomboentity.UpdateFlexiComboResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
