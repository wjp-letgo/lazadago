package sellervoucher

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	sellervoucherentity "github.com/wjpxxx/lazadago/sellervoucher/entity"
)
//SellerVoucher
type SellerVoucher struct{
	Config *lazadaConfig.Config
}

//SellerVoucherActivate
//@Title activate seller voucher promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/activate
func (s *SellerVoucher) SellerVoucherActivate (voucherType string,id int64) sellervoucherentity.SellerVoucherActivateResult {
    method := "/promotion/voucher/activate"
    params := lib.InRow{
      "voucher_type":voucherType,
      "id":id,
    }
    result := sellervoucherentity.SellerVoucherActivateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherAddSelectedProductSKU
//@Title add seller voucher promotion product sku
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/product/sku/add
func (s *SellerVoucher) SellerVoucherAddSelectedProductSKU (voucherType string,id int64,skuIds []int64) sellervoucherentity.SellerVoucherAddSelectedProductSKUResult {
    method := "/promotion/voucher/product/sku/add"
    params := lib.InRow{
      "voucher_type":voucherType,
      "id":id,
      "sku_ids":skuIds,
    }
    result := sellervoucherentity.SellerVoucherAddSelectedProductSKUResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherCreate
//@Title create a new seller voucher promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/create
func (s *SellerVoucher) SellerVoucherCreate (criteriaOverMoney string,voucherType string,apply string,collectStart int,displayArea string,periodEndTime int,voucherName string,voucherDiscountType string,offeringMoneyValueOff string,periodStartTime int,limit int,issued int,maxDiscountOfferingMoneyValue string,offeringPercentageDiscountOff int) sellervoucherentity.SellerVoucherCreateResult {
    method := "/promotion/voucher/create"
    params := lib.InRow{
      "criteria_over_money":criteriaOverMoney,
      "voucher_type":voucherType,
      "apply":apply,
      "display_area":displayArea,
      "period_end_time":periodEndTime,
      "voucher_name":voucherName,
      "voucher_discount_type":voucherDiscountType,
      "period_start_time":periodStartTime,
      "limit":limit,
      "issued":issued,
    }
    if collectStart>-1{
        params["collect_start"]=collectStart
    }
    if offeringMoneyValueOff!=""{
        params["offering_money_value_off"]=offeringMoneyValueOff
    }
    if maxDiscountOfferingMoneyValue!=""{
        params["max_discount_offering_money_value"]=maxDiscountOfferingMoneyValue
    }
    if offeringPercentageDiscountOff>-1{
        params["offering_percentage_discount_off"]=offeringPercentageDiscountOff
    }
    result := sellervoucherentity.SellerVoucherCreateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherDeactivate
//@Title deactivate seller voucher promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/deactivate
func (s *SellerVoucher) SellerVoucherDeactivate (voucherType string,id int64) sellervoucherentity.SellerVoucherDeactivateResult {
    method := "/promotion/voucher/deactivate"
    params := lib.InRow{
      "voucher_type":voucherType,
      "id":id,
    }
    result := sellervoucherentity.SellerVoucherDeactivateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherDeleteSelectedProductSKU
//@Title delete seller voucher promotion product sku
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/product/sku/remove
func (s *SellerVoucher) SellerVoucherDeleteSelectedProductSKU (voucherType string,id int64,skuIds []int64) sellervoucherentity.SellerVoucherDeleteSelectedProductSKUResult {
    method := "/promotion/voucher/product/sku/remove"
    params := lib.InRow{
      "voucher_type":voucherType,
      "id":id,
      "sku_ids":skuIds,
    }
    result := sellervoucherentity.SellerVoucherDeleteSelectedProductSKUResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherDetailQuery
//@Title get a seller voucher promotion detail
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/get
func (s *SellerVoucher) SellerVoucherDetailQuery (voucherType string,id int64) sellervoucherentity.SellerVoucherDetailQueryResult {
    method := "/promotion/voucher/get"
    params := lib.InRow{
      "voucher_type":voucherType,
      "id":id,
    }
    result := sellervoucherentity.SellerVoucherDetailQueryResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherList
//@Title query seller voucher promotion list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/vouchers/get
func (s *SellerVoucher) SellerVoucherList (curPage int,voucherType string,name string,pageSize int,status string) sellervoucherentity.SellerVoucherListResult {
    method := "/promotion/vouchers/get"
    params := lib.InRow{
      "voucher_type":voucherType,
    }
    if curPage>-1{
        params["cur_page"]=curPage
    }
    if name!=""{
        params["name"]=name
    }
    if pageSize>-1{
        params["page_size"]=pageSize
    }
    if status!=""{
        params["status"]=status
    }
    result := sellervoucherentity.SellerVoucherListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherSelectedProductList
//@Title query seller voucher selected products list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/products/get
func (s *SellerVoucher) SellerVoucherSelectedProductList (voucherType string,id int64) sellervoucherentity.SellerVoucherSelectedProductListResult {
    method := "/promotion/voucher/products/get"
    params := lib.InRow{
      "voucher_type":voucherType,
      "id":id,
    }
    result := sellervoucherentity.SellerVoucherSelectedProductListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SellerVoucherUpdate
//@Title update a existing seller voucher promotion
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=30&path=/promotion/voucher/update
func (s *SellerVoucher) SellerVoucherUpdate (maxDiscountOfferingMoneyValue string,offeringPercentageDiscountOff int,id string,criteriaOverMoney string,voucherType string,apply string,collectStart int,displayArea string,periodEndTime int,voucherName string,voucherDiscountType string,offeringMoneyValueOff string,periodStartTime int,limit int,issued int) sellervoucherentity.SellerVoucherUpdateResult {
    method := "/promotion/voucher/update"
    params := lib.InRow{
      "id":id,
      "criteria_over_money":criteriaOverMoney,
      "voucher_type":voucherType,
      "apply":apply,
      "display_area":displayArea,
      "period_end_time":periodEndTime,
      "voucher_name":voucherName,
      "voucher_discount_type":voucherDiscountType,
      "offering_money_value_off":offeringMoneyValueOff,
      "period_start_time":periodStartTime,
      "limit":limit,
      "issued":issued,
    }
    if maxDiscountOfferingMoneyValue!=""{
        params["max_discount_offering_money_value"]=maxDiscountOfferingMoneyValue
    }
    if offeringPercentageDiscountOff>-1{
        params["offering_percentage_discount_off"]=offeringPercentageDiscountOff
    }
    if collectStart>-1{
        params["collect_start"]=collectStart
    }
    result := sellervoucherentity.SellerVoucherUpdateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
