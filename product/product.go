package product

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	productentity "github.com/wjpxxx/lazadago/product/entity"
    "fmt"
)
//Product
type Product struct{
	Config *lazadaConfig.Config
}

//CreateProduct
//@Title Use this API to create a single new product.
//25/6/20 : Updated for DBS changes. Refer to Input Parameters Payload
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/create
func (s *Product) CreateProduct (payload string) productentity.CreateProductResult {
    method := "/product/create"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.CreateProductResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//DeactivateProduct
//@Title Use this API to deactivate Product or SKUs corresponding to the product
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/deactivate
func (s *Product) DeactivateProduct (apiRequestBody string) productentity.DeactivateProductResult {
    method := "/product/deactivate"
    params := lib.InRow{
      "apiRequestBody":apiRequestBody,
    }
    result := productentity.DeactivateProductResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetBrandByPages
//@Title Use this API to retrieve all product brands by page index in the system.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/category/brands/query
func (s *Product) GetBrandByPagesInt (startRow int,pageSize int,languageCode string) productentity.GetBrandByPagesResult {
    return s.GetBrandByPages(fmt.Sprintf("%d",startRow),fmt.Sprintf("%d",pageSize),languageCode)
}
//GetBrandByPages
//@Title Use this API to retrieve all product brands by page index in the system.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/category/brands/query
func (s *Product) GetBrandByPages (startRow string,pageSize string,languageCode string) productentity.GetBrandByPagesResult {
    method := "/category/brands/query"
    params := lib.InRow{
      "startRow":startRow,
      "pageSize":pageSize,
      "languageCode":languageCode,
    }
    result := productentity.GetBrandByPagesResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetCategoryAttributes
//@Title Use this API to get a list of attributes for a specified product category.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/category/attributes/get
func (s *Product) GetCategoryAttributesInt64 (primaryCategoryId int64,languageCode string) productentity.GetCategoryAttributesResult {
    return s.GetCategoryAttributes(fmt.Sprintf("%d",primaryCategoryId),languageCode)
}
//GetCategoryAttributes
//@Title Use this API to get a list of attributes for a specified product category.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/category/attributes/get
func (s *Product) GetCategoryAttributes (primaryCategoryId string,languageCode string) productentity.GetCategoryAttributesResult {
    method := "/category/attributes/get"
    params := lib.InRow{
      "primary_category_id":primaryCategoryId,
    }
    if languageCode!=""{
        params["language_code"]=languageCode
    }
    result := productentity.GetCategoryAttributesResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetCategorySuggestion
//@Title Get product's category suggestion by product title
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/category/suggestion/get
func (s *Product) GetCategorySuggestion (productName string) productentity.GetCategorySuggestionResult {
    method := "/product/category/suggestion/get"
    params := lib.InRow{
      "product_name":productName,
    }
    result := productentity.GetCategorySuggestionResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetCategoryTree
//@Title Use this API to retrieve the list of all product categories in the system.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/category/tree/get
func (s *Product) GetCategoryTree (languageCode string) productentity.GetCategoryTreeResult {
    method := "/category/tree/get"
    params := lib.InRow{
    }
    if languageCode!=""{
        params["language_code"]=languageCode
    }
    result := productentity.GetCategoryTreeResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetProductItem
//@Title Get single product by ItemId or SellerSku.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/item/get
func (s *Product) GetProductItem (itemId int64,sellerSku string) productentity.GetProductItemResult {
    method := "/product/item/get"
    params := lib.InRow{
    }
    if itemId>-1{
        params["item_id"]=itemId
    }
    if sellerSku!=""{
        params["seller_sku"]=sellerSku
    }
    result := productentity.GetProductItemResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetProducts
//@Title Use this API to get detailed information of the specified products.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/products/get
func (s *Product) GetProductsByUpdateTime (start,end int,filter string,offset int,limit int) productentity.GetProductsResult {
    updateBeforeStr:=lib.TimeToFormat(end,"2006-01-02T15:04:05+0800")
    updateAfterStr:=lib.TimeToFormat(start,"2006-01-02T15:04:05+0800")
    return s.GetProducts(filter,updateBeforeStr,"",fmt.Sprintf("%d",offset) ,"",updateAfterStr,fmt.Sprintf("%d",limit),"","")
}
//GetProducts
//@Title Use this API to get detailed information of the specified products.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/products/get
func (s *Product) GetProducts (filter string,updateBefore string,createBefore string,offset string,createAfter string,updateAfter string,limit string,options string,skuSellerList string) productentity.GetProductsResult {
    method := "/products/get"
    params := lib.InRow{
      "filter":filter,
    }
    if updateBefore!=""{
        params["update_before"]=updateBefore
    }
    if createBefore!=""{
        params["create_before"]=createBefore
    }
    if offset!=""{
        params["offset"]=offset
    }
    if createAfter!=""{
        params["create_after"]=createAfter
    }
    if updateAfter!=""{
        params["update_after"]=updateAfter
    }
    if limit!=""{
        params["limit"]=limit
    }
    if options!=""{
        params["options"]=options
    }
    if skuSellerList!=""{
        params["sku_seller_list"]=skuSellerList
    }
    result := productentity.GetProductsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetQcStatus
//@Title Use this API to get the quality control status of items being listed.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/qc/status/get
func (s *Product) GetQcStatus (offset int,limit int,sellerSkus []string) productentity.GetQcStatusResult {
    method := "/product/qc/status/get"
    params := lib.InRow{
      "offset":offset,
      "limit":limit,
      "seller_skus":sellerSkus,
    }
    result := productentity.GetQcStatusResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetResponse
//@Title Use this API to get the returned information from the system for the MigrateImages API.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/image/response/get
func (s *Product) GetResponse (batchId string) productentity.GetResponseResult {
    method := "/image/response/get"
    params := lib.InRow{
      "batch_id":batchId,
    }
    result := productentity.GetResponseResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetSellerItemLimit
//@Title The platform will provide the product quantity limit information by this interface. The qps will be limited by seller, 10 qps per seller.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/seller/item/limit
func (s *Product) GetSellerItemLimit () productentity.GetSellerItemLimitResult {
    method := "/product/seller/item/limit"
    params := lib.InRow{}
    result := productentity.GetSellerItemLimitResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetUnfilledAttributeItem
//@Title Get products without key attributes
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/unfilled/attribute/get
func (s *Product) GetUnfilledAttributeItem (pageIndex int,attributeTag string,pageSize int,languageCode string) productentity.GetUnfilledAttributeItemResult {
    method := "/product/unfilled/attribute/get"
    params := lib.InRow{
      "page_index":pageIndex,
      "attribute_tag":attributeTag,
      "page_size":pageSize,
      "language_code":languageCode,
    }
    result := productentity.GetUnfilledAttributeItemResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//MigrateImage
//@Title distributor purchase orders query
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/image/migrate
func (s *Product) MigrateImage (payload string) productentity.MigrateImageResult {
    method := "/image/migrate"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.MigrateImageResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//MigrateImages
//@Title Use this API to migrate multiple images from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB. A single call can migrate 8 images at most.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/images/migrate
func (s *Product) MigrateImages (payload string) productentity.MigrateImagesResult {
    method := "/images/migrate"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.MigrateImagesResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//RemoveProduct
//@Title Use this API to remove an existing product, some SKUs in one product, or all SKUs in one product. System supports a maximum number of 50 SellerSkus in one request.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/remove
func (s *Product) RemoveProduct (sellerSkuList string) productentity.RemoveProductResult {
    method := "/product/remove"
    params := lib.InRow{
      "seller_sku_list":sellerSkuList,
    }
    result := productentity.RemoveProductResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//RemoveSku
//@Title Use this API to delete SKUs and sales attributes of corresponding products.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/sku/remove
func (s *Product) RemoveSku (payload string) productentity.RemoveSkuResult {
    method := "/product/sku/remove"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.RemoveSkuResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetImages
//@Title Use this API to set the images for an existing product by associating one or more image URLs with it.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/images/set
func (s *Product) SetImages (payload string) productentity.SetImagesResult {
    method := "/images/set"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.SetImagesResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UpdatePriceQuantity
//@Title Use this API to update the price and quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/price_quantity/update
func (s *Product) UpdatePriceQuantity (payload string) productentity.UpdatePriceQuantityResult {
    method := "/product/price_quantity/update"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.UpdatePriceQuantityResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UpdateProduct
//@Title Use this API to update attributes or SKUs of an existing product.
//The iteration 25/6/2020 Updated for DBS changes. Refer to Input Parameters Payload
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/update
func (s *Product) UpdateProduct (payload string) productentity.UpdateProductResult {
    method := "/product/update"
    params := lib.InRow{
      "payload":payload,
    }
    result := productentity.UpdateProductResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UploadImage
//@Title Use this API to upload a single image file to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/image/upload
func (s *Product) UploadImage (image []byte) productentity.UploadImageResult {
    method := "/image/upload"
    params := lib.InRow{
      "@image":image,
    }
    result := productentity.UploadImageResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UploadImage
//@Title Use this API to upload a single image file to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/image/upload
func (s *Product) UploadImageByPath (imagePath string) productentity.UploadImageResult {
    method := "/image/upload"
    params := lib.InRow{
      "@image":imagePath,
    }
    result := productentity.UploadImageResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//RetailFulfilmentCreate
//@Title fulfilment create
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/fulfillment/create
func (s *Product) RetailFulfilmentCreate (platformName string,source string,sellerId int64,platformSkuCode string,itemId int64,skuId int64,platformSkuName string,barcodeList []string,categoryId int64,brand string,brandName string,isShelfLifeMgt bool,lifeCycleDays int,rejectLifeCycleDays int,lockupLifeCycleDays int,adventLifeCycleDays int,isSnMgt bool,cpWeight int,cpLength int,cpWidth int,cpHeight int,skuPrice int,features productentity.RetailFulfilmentCreateFeaturesRequestEntity) productentity.RetailFulfilmentCreateResult {
    method := "/product/fulfillment/create"
    params := lib.InRow{
      "platform_name":platformName,
      "source":source,
      "seller_id":sellerId,
      "platform_sku_code":platformSkuCode,
      "item_id":itemId,
      "sku_id":skuId,
      "platform_sku_name":platformSkuName,
      "barcode_list":barcodeList,
      "category_id":categoryId,
      "brand_name":brandName,
      "features":features,
    }
    if brand!=""{
        params["brand"]=brand
    }
    if lifeCycleDays>-1{
        params["life_cycle_days"]=lifeCycleDays
    }
    if rejectLifeCycleDays>-1{
        params["reject_life_cycle_days"]=rejectLifeCycleDays
    }
    if lockupLifeCycleDays>-1{
        params["lockup_life_cycle_days"]=lockupLifeCycleDays
    }
    if adventLifeCycleDays>-1{
        params["advent_life_cycle_days"]=adventLifeCycleDays
    }
    if cpWeight>-1{
        params["cp_weight"]=cpWeight
    }
    if cpLength>-1{
        params["cp_length"]=cpLength
    }
    if cpWidth>-1{
        params["cp_width"]=cpWidth
    }
    if cpHeight>-1{
        params["cp_height"]=cpHeight
    }
    if skuPrice>-1{
        params["sku_price"]=skuPrice
    }
    result := productentity.RetailFulfilmentCreateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//RetailFulfilmentUpdate
//@Title fulfilment update
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=5&path=/product/fulfillment/update
func (s *Product) RetailFulfilmentUpdate (scItemId string,fulfillmentSkuName string,barcodeList []string,categoryId int64,brand string,brandName string,isShelfLifeMgt bool,lifeCycleDays int,rejectLifeCycleDays int,lockupLifeCycleDays int,adventLifeCycleDays int,isSnMgt bool,cpWeight int,cpLength int,cpWidth int,cpHeight int,skuPrice int,features productentity.RetailFulfilmentUpdateFeaturesRequestEntity,source string) productentity.RetailFulfilmentUpdateResult {
    method := "/product/fulfillment/update"
    params := lib.InRow{
      "sc_item_id":scItemId,
    }
    if fulfillmentSkuName!=""{
        params["fulfillment_sku_name"]=fulfillmentSkuName
    }
    if categoryId>-1{
        params["category_id"]=categoryId
    }
    if brand!=""{
        params["brand"]=brand
    }
    if brandName!=""{
        params["brand_name"]=brandName
    }
    if lifeCycleDays>-1{
        params["life_cycle_days"]=lifeCycleDays
    }
    if rejectLifeCycleDays>-1{
        params["reject_life_cycle_days"]=rejectLifeCycleDays
    }
    if lockupLifeCycleDays>-1{
        params["lockup_life_cycle_days"]=lockupLifeCycleDays
    }
    if adventLifeCycleDays>-1{
        params["advent_life_cycle_days"]=adventLifeCycleDays
    }
    if cpWeight>-1{
        params["cp_weight"]=cpWeight
    }
    if cpLength>-1{
        params["cp_length"]=cpLength
    }
    if cpWidth>-1{
        params["cp_width"]=cpWidth
    }
    if cpHeight>-1{
        params["cp_height"]=cpHeight
    }
    if skuPrice>-1{
        params["sku_price"]=skuPrice
    }
    if source!=""{
        params["source"]=source
    }
    result := productentity.RetailFulfilmentUpdateResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
