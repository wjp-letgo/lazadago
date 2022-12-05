package lazadago

import (
	"github.com/wjpxxx/lazadago/auth"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	"github.com/wjpxxx/lazadago/system"
	systementity "github.com/wjpxxx/lazadago/system/entity"
	"github.com/wjpxxx/lazadago/order"
	orderentity "github.com/wjpxxx/lazadago/order/entity"
	"github.com/wjpxxx/lazadago/product"
	productentity "github.com/wjpxxx/lazadago/product/entity"
    "github.com/wjpxxx/lazadago/finance"
	financeentity "github.com/wjpxxx/lazadago/finance/entity"
    "github.com/wjpxxx/lazadago/logistics"
	logisticsentity "github.com/wjpxxx/lazadago/logistics/entity"
    "github.com/wjpxxx/lazadago/seller"
	sellerentity "github.com/wjpxxx/lazadago/seller/entity"
    "github.com/wjpxxx/lazadago/datamoat"
	datamoatentity "github.com/wjpxxx/lazadago/datamoat/entity"
    "github.com/wjpxxx/lazadago/fbl"
	fblentity "github.com/wjpxxx/lazadago/fbl/entity"
    "github.com/wjpxxx/lazadago/gspproduct"
	gspproductentity "github.com/wjpxxx/lazadago/gspproduct/entity"
    "github.com/wjpxxx/lazadago/etickets"
	eticketsentity "github.com/wjpxxx/lazadago/etickets/entity"
    "github.com/wjpxxx/lazadago/returnandrefund"
	returnandrefundentity "github.com/wjpxxx/lazadago/returnandrefund/entity"
    "github.com/wjpxxx/lazadago/flexicombo"
	flexicomboentity "github.com/wjpxxx/lazadago/flexicombo/entity"
    "github.com/wjpxxx/lazadago/sellervoucher"
	sellervoucherentity "github.com/wjpxxx/lazadago/sellervoucher/entity"
    "github.com/wjpxxx/lazadago/freeshipping"
	freeshippingentity "github.com/wjpxxx/lazadago/freeshipping/entity"
    "github.com/wjpxxx/lazadago/mediacenter"
	mediacenterentity "github.com/wjpxxx/lazadago/mediacenter/entity"
    "github.com/wjpxxx/lazadago/productreview"
	productreviewentity "github.com/wjpxxx/lazadago/productreview/entity"
    "github.com/wjpxxx/lazadago/firstmilebigbag"
	firstmilebigbagentity "github.com/wjpxxx/lazadago/firstmilebigbag/entity"
)

//Lazadar
type Lazadar interface {
    SetAccessToken(accessToken string)
	//auth
	AuthorizationURL(redirectUri string) string
	//system
	GenerateAccessToken (code string,uuid string) systementity.GenerateAccessTokenResult
	RefreshAccessToken (refreshToken string) systementity.RefreshAccessTokenResult
	//order
    GetOrdersByUpdateTime (start,end int,offset int,limit int, status string) orderentity.GetOrdersResult
	GetAwbDocumentHtml (orderItemIds string) orderentity.GetAwbDocumentHtmlResult
    GetAwbDocumentPDF (orderItemIds string) orderentity.GetAwbDocumentPDFResult     
    GetDocument (docType string,orderItemIds string) orderentity.GetDocumentResult     
    GetFailureReasons () orderentity.GetFailureReasonsResult     
    GetMultipleOrderItems (orderIds []int64) orderentity.GetMultipleOrderItemsResult     
    GetOVOOrders (tradeOrderIds string) orderentity.GetOVOOrdersResult     
    GetOrder (orderId int64) orderentity.GetOrderResult     
    GetOrderItems (orderId int64) orderentity.GetOrderItemsResult     
    GetOrders (updateBefore string,sortDirection string,offset int,limit int,updateAfter string,sortBy string,createdBefore string,createdAfter string,status string) orderentity.GetOrdersResult     
    SetInvoiceNumber (orderItemId int64,invoiceNumber string) orderentity.SetInvoiceNumberResult     
    SetRepack (packageId string) orderentity.SetRepackResult     
    SetStatusToCanceled (reasonDetail string,reasonId int64,orderItemId int64) orderentity.SetStatusToCanceledResult     
    SetStatusToPackedByMarketplace (shippingProvider string,deliveryType string,orderItemIds string) orderentity.SetStatusToPackedByMarketplaceResult     
    SetStatusToReadyToShip (deliveryType string,orderItemIds string,shipmentProvider string,trackingNumber string) orderentity.SetStatusToReadyToShipResult     
    SetStatusToSOFDelivered (orderItemIds string) orderentity.SetStatusToSOFDeliveredResult     
    SetStatusToSOFFailedDelivery (orderItemIds string) orderentity.SetStatusToSOFFailedDeliveryResult
	//product
    GetProductsByUpdateTime (start,end int,filter string,offset int,limit int) productentity.GetProductsResult
    GetBrandByPagesInt (startRow int,pageSize int,languageCode string) productentity.GetBrandByPagesResult
    GetCategoryAttributesInt64 (primaryCategoryId int64,languageCode string) productentity.GetCategoryAttributesResult
    UploadImageByPath (imagePath string) productentity.UploadImageResult
    CreateProduct (payload string) productentity.CreateProductResult
    DeactivateProduct (apiRequestBody string) productentity.DeactivateProductResult
    GetBrandByPages (startRow string,pageSize string,languageCode string) productentity.GetBrandByPagesResult
    GetCategoryAttributes (primaryCategoryId string,languageCode string) productentity.GetCategoryAttributesResult
    GetCategorySuggestion (productName string) productentity.GetCategorySuggestionResult
    GetCategoryTree (languageCode string) productentity.GetCategoryTreeResult
    GetProductItem (itemId int64,sellerSku string) productentity.GetProductItemResult
    GetProducts (filter string,updateBefore string,createBefore string,offset string,createAfter string,updateAfter string,limit string,options string,skuSellerList string) productentity.GetProductsResult
    GetQcStatus (offset int,limit int,sellerSkus []string) productentity.GetQcStatusResult
    GetResponse (batchId string) productentity.GetResponseResult
    GetSellerItemLimit () productentity.GetSellerItemLimitResult
    GetUnfilledAttributeItem (pageIndex int,attributeTag string,pageSize int,languageCode string) productentity.GetUnfilledAttributeItemResult
    MigrateImage (payload string) productentity.MigrateImageResult
    MigrateImages (payload string) productentity.MigrateImagesResult
    RemoveProduct (sellerSkuList string) productentity.RemoveProductResult
    RemoveSku (payload string) productentity.RemoveSkuResult
    SetImages (payload string) productentity.SetImagesResult
    UpdatePriceQuantity (payload string) productentity.UpdatePriceQuantityResult
    UpdateProduct (payload string) productentity.UpdateProductResult
    UploadImage (image []byte) productentity.UploadImageResult
    RetailFulfilmentCreate (platformName string,source string,sellerId int64,platformSkuCode string,itemId int64,skuId int64,platformSkuName string,barcodeList []string,categoryId int64,brand string,brandName string,isShelfLifeMgt bool,lifeCycleDays int,rejectLifeCycleDays int,lockupLifeCycleDays int,adventLifeCycleDays int,isSnMgt bool,cpWeight int,cpLength int,cpWidth int,cpHeight int,skuPrice int,features productentity.RetailFulfilmentCreateFeaturesRequestEntity) productentity.RetailFulfilmentCreateResult
    RetailFulfilmentUpdate (scItemId string,fulfillmentSkuName string,barcodeList []string,categoryId int64,brand string,brandName string,isShelfLifeMgt bool,lifeCycleDays int,rejectLifeCycleDays int,lockupLifeCycleDays int,adventLifeCycleDays int,isSnMgt bool,cpWeight int,cpLength int,cpWidth int,cpHeight int,skuPrice int,features productentity.RetailFulfilmentUpdateFeaturesRequestEntity,source string) productentity.RetailFulfilmentUpdateResult
	//Finance
    GetPayoutStatus (createdAfter string) financeentity.GetPayoutStatusResult
    GetTransactionDetails (transType string,startTime string,endTime string,limit string,offset string) financeentity.GetTransactionDetailsResult
    QueryTransactionDetails (offset string,transType string,tradeOrderId string,limit string,startTime string,endTime string,tradeOrderLineId string) financeentity.QueryTransactionDetailsResult
    //Logistics
    GetOrderTrace (sellerId string,orderId string,locale string,ofcPackageIdList []string) logisticsentity.GetOrderTraceResult
    GetShipmentProviders () logisticsentity.GetShipmentProvidersResult
    //Seller
    GetMultiWarehouseBySeller (addressTypes []string) sellerentity.GetMultiWarehouseBySellerResult
    GetSeller () sellerentity.GetSellerResult
    GetSellerMetricsById () sellerentity.GetSellerMetricsByIdResult
    GetSellerPerformance (language string) sellerentity.GetSellerPerformanceResult
    UpdateSeller (payload string) sellerentity.UpdateSellerResult
    UpdateUser (payload string) sellerentity.UpdateUserResult
    //DataMoat
    DataMoatComputeRisk (time string,appName string,userId string,userIp string,ati string) datamoatentity.DataMoatComputeRiskResult
    DataMoatLogin (time string,appName string,userId string,tid string,userIp string,ati string,loginResult string,loginMessage string) datamoatentity.DataMoatLoginResult
    //FBL
    GetPlatformProducts (perPage int,sellerId int64,marketplace string,sellerSku string,platformSkuName string,readyForInbound bool,platformSku string,page int) fblentity.GetPlatformProductsResult
    GetFulfillmentProductDetail (perPage int,shelfLifeFlag bool,marketplace string,fulfillmentSku string,serialNumberFlag bool,page int,fulfillmentSkuName string,barcode string) fblentity.GetFulfillmentProductDetailResult
    GetInboundOrderDetail (inboundOrderNo string,marketplace string) fblentity.GetInboundOrderDetailResult
    GetInboundOrderList (inboundOrderNo string,creationTimeFrom string,creationTimeTo string,inboundWarehouse string,sellerSku string,fulfillmentSku string,marketplace string,page string,perPage string) fblentity.GetInboundOrderListResult
    GetInventoryChangedSKU (warehouseCode string,page int,perPage int,marketPlace string,operateTimeFrom string,operateTimeTo string) fblentity.GetInventoryChangedSKUResult
    GetInventoryOperateLog (page int,perPage int,marketPlace string,operateTimeFrom string,operateTimeTo string,warehouseCode string,fulfillmentSkuId string) fblentity.GetInventoryOperateLogResult
    GetOutboundOrderDetail (outboundOrderNo string,marketplace string) fblentity.GetOutboundOrderDetailResult
    GetOutboundOrderList (outboundOrderNo string,creationTimeFrom string,creationTimeTo string,outboundWarehouse string,sellerSku string,fulfillmentSku string,marketplace string,page string,perPage string) fblentity.GetOutboundOrderListResult
    GetWarehouseStock (sellerSku string,marketplace string,fulfilmentSku string,storeCode string) fblentity.GetWarehouseStockResult
    GetWarehouseStockV3 (sellerSku string,marketplace string,fulfilmentSku string,storeCode string) fblentity.GetWarehouseStockV3Result
    UploadWaybill (waybill []byte,packageCode string,trackingNumber string,extendsField string,storeCode string) fblentity.UploadWaybillResult
    //GSPProduct API
    CreateGlobalProduct (payload string) gspproductentity.CreateGlobalProductResult
    GetGlobalProductStatus (params gspproductentity.GetGlobalProductStatusParamsRequestEntity) gspproductentity.GetGlobalProductStatusResult
    GetUnfilledAttribute (offset int,limit int,attributeTag string) gspproductentity.GetUnfilledAttributeResult
    UpdateGlobalProductAttribute (payload string) gspproductentity.UpdateGlobalProductAttributeResult
    //E-Tickets API
    GetOrderItemsFromBarCode (code string) eticketsentity.GetOrderItemsFromBarCodeResult
    RedeemOrderItems (bizType int,code string,outerId string,serialNum string,consumeNum int,storeId string,posId string) eticketsentity.RedeemOrderItemsResult
    GlobalEticketMerchantMaAvailable (bizType int,code string,serialNum string,posId string,outerId string,consumeNum int,consumeStoreId string) eticketsentity.GlobalEticketMerchantMaAvailableResult
    GlobalEticketMerchantMaConsume (bizType int,serialNum string,posId string,outerId string,consumeNum int,code string,consumeStoreId string) eticketsentity.GlobalEticketMerchantMaConsumeResult
    GlobalEticketMerchantMaFailsend (bizType int,subCode string,outerId string,subMsg string) eticketsentity.GlobalEticketMerchantMaFailsendResult
    GlobalEticketMerchantMaQuery (code string,sellerId int64,storeId int64) eticketsentity.GlobalEticketMerchantMaQueryResult
    GlobalEticketMerchantMaQueryTbMa (code string) eticketsentity.GlobalEticketMerchantMaQueryTbMaResult
    GlobalEticketMerchantMaSend (bizType int,isvMaList []eticketsentity.GlobalEticketMerchantMaSendIsvMaListRequestEntity,outerId string) eticketsentity.GlobalEticketMerchantMaSendResult
    //Return and Refund API
    GetReverseOrderDetail (reverseOrderId int64) returnandrefundentity.GetReverseOrderDetailResult
    GetReverseOrderHistoryList (reverseOrderLineId int64,pageSize int,pageNumber int) returnandrefundentity.GetReverseOrderHistoryListResult
    GetReverseOrderReasonList (reverseOrderLineId int64) returnandrefundentity.GetReverseOrderReasonListResult
    GetReverseOrdersForSeller (ofcStatusList []string,reverseOrderId int64,tradeOrderId int64,pageSize int,reverseStatusList []string,pageNo int,returnToType string,disputeInProgress bool) returnandrefundentity.GetReverseOrdersForSellerResult
    InitReverseOrderCancel (orderItemIdList []string,orderId int64,reasonId string) returnandrefundentity.InitReverseOrderCancelResult
    ReverseOrderCancelValidate (orderId string,orderItemIdList []string) returnandrefundentity.ReverseOrderCancelValidateResult
    ReverseOrderReturnUpdate (action string,reverseOrderId int64,reverseOrderItemIds []int64,reasonId int64,comment string,imageInfo []returnandrefundentity.ReverseOrderReturnUpdateImageInfoRequestEntity) returnandrefundentity.ReverseOrderReturnUpdateResult
    //FlexiCombo
    ActivateFlexiCombo (id int64) flexicomboentity.ActivateFlexiComboResult
    AddFlexiComboProducts (id int64,skuIds []int64) flexicomboentity.AddFlexiComboProductsResult
    CreateFlexiCombo (apply string,sampleSkus []flexicomboentity.CreateFlexiComboSampleSkusRequestEntity,criteriaType string,criteriaValue []string,orderNumbers int,name string,platformChannel string,giftSkus []flexicomboentity.CreateFlexiComboGiftSkusRequestEntity,startTime int,discountType string,endTime int,discountValue []string,stackable string) flexicomboentity.CreateFlexiComboResult
    DeactivateFlexiCombo (id int64) flexicomboentity.DeactivateFlexiComboResult
    DeleteFlexiComboProducts (id int64,skuIds []int64) flexicomboentity.DeleteFlexiComboProductsResult
    GetFlexiComboDetails (id int64) flexicomboentity.GetFlexiComboDetailsResult
    ListFlexiCombo (curPage int,name string,pageSize int,status string) flexicomboentity.ListFlexiComboResult
    ListFlexiComboProducts (curPage int,pageSize int,id int64) flexicomboentity.ListFlexiComboProductsResult
    UpdateFlexiCombo (apply string,sampleSkus []flexicomboentity.UpdateFlexiComboSampleSkusRequestEntity,criteriaType string,criteriaValue []string,orderNumbers int,name string,platformChannel string,giftSkus []flexicomboentity.UpdateFlexiComboGiftSkusRequestEntity,startTime int,discountType string,id int64,endTime int,discountValue []string,stackable string) flexicomboentity.UpdateFlexiComboResult
    //Seller Voucher
    SellerVoucherActivate (voucherType string,id int64) sellervoucherentity.SellerVoucherActivateResult
    SellerVoucherAddSelectedProductSKU (voucherType string,id int64,skuIds []int64) sellervoucherentity.SellerVoucherAddSelectedProductSKUResult
    SellerVoucherCreate (criteriaOverMoney string,voucherType string,apply string,collectStart int,displayArea string,periodEndTime int,voucherName string,voucherDiscountType string,offeringMoneyValueOff string,periodStartTime int,limit int,issued int,maxDiscountOfferingMoneyValue string,offeringPercentageDiscountOff int) sellervoucherentity.SellerVoucherCreateResult
    SellerVoucherDeactivate (voucherType string,id int64) sellervoucherentity.SellerVoucherDeactivateResult
    SellerVoucherDeleteSelectedProductSKU (voucherType string,id int64,skuIds []int64) sellervoucherentity.SellerVoucherDeleteSelectedProductSKUResult
    SellerVoucherDetailQuery (voucherType string,id int64) sellervoucherentity.SellerVoucherDetailQueryResult
    SellerVoucherList (curPage int,voucherType string,name string,pageSize int,status string) sellervoucherentity.SellerVoucherListResult
    SellerVoucherSelectedProductList (voucherType string,id int64) sellervoucherentity.SellerVoucherSelectedProductListResult
    SellerVoucherUpdate (maxDiscountOfferingMoneyValue string,offeringPercentageDiscountOff int,id string,criteriaOverMoney string,voucherType string,apply string,collectStart int,displayArea string,periodEndTime int,voucherName string,voucherDiscountType string,offeringMoneyValueOff string,periodStartTime int,limit int,issued int) sellervoucherentity.SellerVoucherUpdateResult
    //Free Shipping
    FreeShippingActivate (id int64) freeshippingentity.FreeShippingActivateResult
    FreeShippingAddSelectedProductSKU (id int64,skuIds []int64) freeshippingentity.FreeShippingAddSelectedProductSKUResult
    FreeShippingCreate (budgetType string,templateType string,apply string,periodEndTime int,templateCode string,categoryName string,budgetValue string,promotionName string,periodType string,regionType string,periodStartTime int,platformChannel string,campaignTag string,regionValue []string,deliveryOption string,tiers []freeshippingentity.FreeShippingCreateTiersRequestEntity,discountType string,dealCriteria string) freeshippingentity.FreeShippingCreateResult
    FreeShippingDeactivate (id int64) freeshippingentity.FreeShippingDeactivateResult
    FreeShippingDeleteSelectedProductSKU (id int64,skuIds []int64) freeshippingentity.FreeShippingDeleteSelectedProductSKUResult
    FreeShippingDeliveryOptionsQuery () freeshippingentity.FreeShippingDeliveryOptionsQueryResult
    FreeShippingGet (id int64) freeshippingentity.FreeShippingGetResult
    FreeShippingList (curPage int,name string,pageSize int,status string) freeshippingentity.FreeShippingListResult
    FreeShippingRegionsQuery () freeshippingentity.FreeShippingRegionsQueryResult
    FreeShippingSelectedProductList (curPage int,pageSize int,id int64) freeshippingentity.FreeShippingSelectedProductListResult
    FreeShippingUpdate (budgetType string,templateType string,apply string,periodEndTime int,templateCode string,categoryName string,budgetValue string,promotionName string,periodType string,regionType string,periodStartTime int,platformChannel string,campaignTag string,regionValue []string,id int64,deliveryOption string,tiers []freeshippingentity.FreeShippingUpdateTiersRequestEntity,discountType string,dealCriteria string) freeshippingentity.FreeShippingUpdateResult
    //Media Center API
    CompleteCreateVideo (uploadId string,parts string,title string,coverUrl string) mediacenterentity.CompleteCreateVideoResult
    GetVideo (videoId int64) mediacenterentity.GetVideoResult
    GetVideoQuota () mediacenterentity.GetVideoQuotaResult
    InitCreateVideo (fileName string,fileBytes int) mediacenterentity.InitCreateVideoResult
    RemoveVideo (videoId int64) mediacenterentity.RemoveVideoResult
    UploadVideoBlock (uploadId string,blockNo string,blockCount string,file []byte) mediacenterentity.UploadVideoBlockResult
    //Product Review API
    GetProductReviewList (itemId int64,orderId int64,startTime int,endTime int,contentFilter string,statusFilter string,pageSize int,current int) productreviewentity.GetProductReviewListResult
    SubmitSellerReply (id int64,content string) productreviewentity.SubmitSellerReplyResult
    //FirstMile Bigbag（only for CN）
    GetLazadaBigbagPDFLable (userInfo firstmilebigbagentity.GetLazadaBigbagPDFLableUserInfoRequestEntity,client string,orderCode string,remark string,locale string,trackingNumber string) firstmilebigbagentity.GetLazadaBigbagPDFLableResult
    GetNewBagNumber () firstmilebigbagentity.GetNewBagNumberResult
    LazadaBigbagCancel (userInfo firstmilebigbagentity.LazadaBigbagCancelUserInfoRequestEntity,client string,orderCode string,remark string,locale string,trackingNumber string) firstmilebigbagentity.LazadaBigbagCancelResult
    LazadaBigbagCollectionPoints (pageSize string,currentPage string) firstmilebigbagentity.LazadaBigbagCollectionPointsResult
    LazadaBigbagCommit (client string,collectionInfo firstmilebigbagentity.LazadaBigbagCommitCollectionInfoRequestEntity,remark string,pickupInfo firstmilebigbagentity.LazadaBigbagCommitPickupInfoRequestEntity,locale string,weightUnit string,tp string,sellerTrackingNumber string,returnInfo firstmilebigbagentity.LazadaBigbagCommitReturnInfoRequestEntity,userInfo firstmilebigbagentity.LazadaBigbagCommitUserInfoRequestEntity,orderCodeList []string,weight string) firstmilebigbagentity.LazadaBigbagCommitResult
    LazadaBigbagUpdate (userInfo firstmilebigbagentity.LazadaBigbagUpdateUserInfoRequestEntity,weight int,locale string,orderCodeList []string,client string,orderCode string,trackingNumber string,weightUnit string) firstmilebigbagentity.LazadaBigbagUpdateResult
    LazadaSellerAccountBind (userInfo firstmilebigbagentity.LazadaSellerAccountBindUserInfoRequestEntity,client string,remark string,sellerList []firstmilebigbagentity.LazadaSellerAccountBindSellerListRequestEntity,locale string) firstmilebigbagentity.LazadaSellerAccountBindResult
    QueryAddressInformaiton (country string,zipCode string,userInfo firstmilebigbagentity.QueryAddressInformaitonUserInfoRequestEntity,city string,remark string,locale string,province string,street string,district string,detailAddress string,client string) firstmilebigbagentity.QueryAddressInformaitonResult
    QueryLazadaBigbagInfo (userInfo firstmilebigbagentity.QueryLazadaBigbagInfoUserInfoRequestEntity,client string,orderCode string,remark string,locale string,trackingNumber string) firstmilebigbagentity.QueryLazadaBigbagInfoResult
    QueryPackageStatus (trackingNumbers []string) firstmilebigbagentity.QueryPackageStatusResult
    UploadLazadaBagNumber (trackingNumbers []string,erpBagNumber string,parcelsTotal int,lzdBagNumber string) firstmilebigbagentity.UploadLazadaBagNumberResult
    UploadSellerBagNumber (trackingNumbers []string,erpBagNumber string,parcelsTotal int) firstmilebigbagentity.UploadSellerBagNumberResult
    GetChannelcodeByFirstMileNo (firstMileNos []string) firstmilebigbagentity.GetChannelcodeByFirstMileNoResult


}

//Lazada
type Lazada struct {
	auth.Auth
	system.System
	order.Order
	product.Product
    finance.Finance
    logistics.Logistics
    seller.Seller
    datamoat.DataMoat
    fbl.Fbl
    gspproduct.GspProduct
    etickets.ETickets
    returnandrefund.ReturnAndRefund
    flexicombo.FlexiCombo
    sellervoucher.SellerVoucher
    freeshipping.FreeShipping
    mediacenter.MediaCenter
    productreview.ProductReview
    firstmilebigbag.FirstMileBigbag
}
//SetAccessToken 设置token
func (l *Lazada)SetAccessToken(accessToken string){
    l.Auth.Config.SetAccessToken(accessToken)
    l.System.Config.SetAccessToken(accessToken)
    l.Order.Config.SetAccessToken(accessToken)
    l.Product.Config.SetAccessToken(accessToken)
    l.Finance.Config.SetAccessToken(accessToken)
    l.Logistics.Config.SetAccessToken(accessToken)
    l.Seller.Config.SetAccessToken(accessToken)
    l.DataMoat.Config.SetAccessToken(accessToken)
    l.Fbl.Config.SetAccessToken(accessToken)
    l.GspProduct.Config.SetAccessToken(accessToken)
    l.ETickets.Config.SetAccessToken(accessToken)
    l.ReturnAndRefund.Config.SetAccessToken(accessToken)
    l.FlexiCombo.Config.SetAccessToken(accessToken)
    l.SellerVoucher.Config.SetAccessToken(accessToken)
    l.FreeShipping.Config.SetAccessToken(accessToken)
    l.MediaCenter.Config.SetAccessToken(accessToken)
    l.ProductReview.Config.SetAccessToken(accessToken)
    l.FirstMileBigbag.Config.SetAccessToken(accessToken)
}

//NewApi
func NewApi(cfg *lazadaConfig.Config)Lazadar{
	return &Lazada{
		auth.Auth{Config: cfg},
		system.System{Config: cfg},
		order.Order{Config: cfg},
		product.Product{Config: cfg},
        finance.Finance{Config:cfg},
        logistics.Logistics{Config: cfg},
        seller.Seller{Config: cfg},
        datamoat.DataMoat{Config: cfg},
        fbl.Fbl{Config: cfg},
        gspproduct.GspProduct{Config: cfg},
        etickets.ETickets{Config: cfg},
        returnandrefund.ReturnAndRefund{Config: cfg},
        flexicombo.FlexiCombo{Config: cfg},
        sellervoucher.SellerVoucher{Config: cfg},
        freeshipping.FreeShipping{Config: cfg},
        mediacenter.MediaCenter{Config: cfg},
        productreview.ProductReview{Config: cfg},
        firstmilebigbag.FirstMileBigbag{Config: cfg},
	}
}
