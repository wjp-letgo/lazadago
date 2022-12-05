package firstmilebigbag

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	firstmilebigbagentity "github.com/wjpxxx/lazadago/firstmilebigbag/entity"
)
//FirstMileBigbag
type FirstMileBigbag struct{
	Config *lazadaConfig.Config
}

//GetLazadaBigbagPDFLable
//@Title Get Lazada Bigbag PDF Lable
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/bigbag/lable/getPdf
func (s *FirstMileBigbag) GetLazadaBigbagPDFLable (userInfo firstmilebigbagentity.GetLazadaBigbagPDFLableUserInfoRequestEntity,client string,orderCode string,remark string,locale string,trackingNumber string) firstmilebigbagentity.GetLazadaBigbagPDFLableResult {
    method := "/logistics/cnpms/bigbag/lable/getPdf"
    params := lib.InRow{
      "userInfo":userInfo,
      "client":client,
    }
    if orderCode!=""{
        params["orderCode"]=orderCode
    }
    if remark!=""{
        params["remark"]=remark
    }
    if locale!=""{
        params["locale"]=locale
    }
    if trackingNumber!=""{
        params["trackingNumber"]=trackingNumber
    }
    result := firstmilebigbagentity.GetLazadaBigbagPDFLableResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetNewBagNumber
//@Title 1.Seller call the api to get the new lzd bag number.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/lzd_bag/get
func (s *FirstMileBigbag) GetNewBagNumber () firstmilebigbagentity.GetNewBagNumberResult {
    method := "/logistics/lzd_bag/get"
    params := lib.InRow{}
    result := firstmilebigbagentity.GetNewBagNumberResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//LazadaBigbagCancel
//@Title Lazada Bigbag cancel
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/bigbag/cancel
func (s *FirstMileBigbag) LazadaBigbagCancel (userInfo firstmilebigbagentity.LazadaBigbagCancelUserInfoRequestEntity,client string,orderCode string,remark string,locale string,trackingNumber string) firstmilebigbagentity.LazadaBigbagCancelResult {
    method := "/logistics/cnpms/bigbag/cancel"
    params := lib.InRow{
      "userInfo":userInfo,
      "client":client,
    }
    if orderCode!=""{
        params["orderCode"]=orderCode
    }
    if remark!=""{
        params["remark"]=remark
    }
    if locale!=""{
        params["locale"]=locale
    }
    if trackingNumber!=""{
        params["trackingNumber"]=trackingNumber
    }
    result := firstmilebigbagentity.LazadaBigbagCancelResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//LazadaBigbagCollectionPoints
//@Title Lazada bigbag query collection points
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/bigbag/querycollection
func (s *FirstMileBigbag) LazadaBigbagCollectionPoints (pageSize string,currentPage string) firstmilebigbagentity.LazadaBigbagCollectionPointsResult {
    method := "/logistics/cnpms/bigbag/querycollection"
    params := lib.InRow{
    }
    if pageSize!=""{
        params["pageSize"]=pageSize
    }
    if currentPage!=""{
        params["currentPage"]=currentPage
    }
    result := firstmilebigbagentity.LazadaBigbagCollectionPointsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//LazadaBigbagCommit
//@Title Lazada bigbag commit
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/bigbag/commit
func (s *FirstMileBigbag) LazadaBigbagCommit (client string,collectionInfo firstmilebigbagentity.LazadaBigbagCommitCollectionInfoRequestEntity,remark string,pickupInfo firstmilebigbagentity.LazadaBigbagCommitPickupInfoRequestEntity,locale string,weightUnit string,tp string,sellerTrackingNumber string,returnInfo firstmilebigbagentity.LazadaBigbagCommitReturnInfoRequestEntity,userInfo firstmilebigbagentity.LazadaBigbagCommitUserInfoRequestEntity,orderCodeList []string,weight string) firstmilebigbagentity.LazadaBigbagCommitResult {
    method := "/logistics/cnpms/bigbag/commit"
    params := lib.InRow{
      "client":client,
      "pickupInfo":pickupInfo,
      "weightUnit":weightUnit,
      "type":tp,
      "returnInfo":returnInfo,
      "userInfo":userInfo,
      "orderCodeList":orderCodeList,
      "weight":weight,
    }
    if remark!=""{
        params["remark"]=remark
    }
    if locale!=""{
        params["locale"]=locale
    }
    if sellerTrackingNumber!=""{
        params["sellerTrackingNumber"]=sellerTrackingNumber
    }
    result := firstmilebigbagentity.LazadaBigbagCommitResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//LazadaBigbagUpdate
//@Title Lazada bigbag update
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/bigbag/update
func (s *FirstMileBigbag) LazadaBigbagUpdate (userInfo firstmilebigbagentity.LazadaBigbagUpdateUserInfoRequestEntity,weight int,locale string,orderCodeList []string,client string,orderCode string,trackingNumber string,weightUnit string) firstmilebigbagentity.LazadaBigbagUpdateResult {
    method := "/logistics/cnpms/bigbag/update"
    params := lib.InRow{
      "userInfo":userInfo,
      "weight":weight,
      "orderCodeList":orderCodeList,
      "client":client,
      "weightUnit":weightUnit,
    }
    if locale!=""{
        params["locale"]=locale
    }
    if orderCode!=""{
        params["orderCode"]=orderCode
    }
    if trackingNumber!=""{
        params["trackingNumber"]=trackingNumber
    }
    result := firstmilebigbagentity.LazadaBigbagUpdateResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//LazadaSellerAccountBind
//@Title Lazada seller account bind for big bag pick up
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/account/bind
func (s *FirstMileBigbag) LazadaSellerAccountBind (userInfo firstmilebigbagentity.LazadaSellerAccountBindUserInfoRequestEntity,client string,remark string,sellerList []firstmilebigbagentity.LazadaSellerAccountBindSellerListRequestEntity,locale string) firstmilebigbagentity.LazadaSellerAccountBindResult {
    method := "/logistics/cnpms/account/bind"
    params := lib.InRow{
      "userInfo":userInfo,
      "sellerList":sellerList,
    }
    if client!=""{
        params["client"]=client
    }
    if remark!=""{
        params["remark"]=remark
    }
    if locale!=""{
        params["locale"]=locale
    }
    result := firstmilebigbagentity.LazadaSellerAccountBindResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//QueryAddressInformaiton
//@Title Query Address Informaiton
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/address/query
func (s *FirstMileBigbag) QueryAddressInformaiton (country string,zipCode string,userInfo firstmilebigbagentity.QueryAddressInformaitonUserInfoRequestEntity,city string,remark string,locale string,province string,street string,district string,detailAddress string,client string) firstmilebigbagentity.QueryAddressInformaitonResult {
    method := "/logistics/cnpms/address/query"
    params := lib.InRow{
      "country":country,
      "userInfo":userInfo,
      "city":city,
      "province":province,
      "street":street,
      "district":district,
      "detailAddress":detailAddress,
    }
    if zipCode!=""{
        params["zipCode"]=zipCode
    }
    if remark!=""{
        params["remark"]=remark
    }
    if locale!=""{
        params["locale"]=locale
    }
    if client!=""{
        params["client"]=client
    }
    result := firstmilebigbagentity.QueryAddressInformaitonResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//QueryLazadaBigbagInfo
//@Title Query Lazada Bigbag Info
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cnpms/bigbag/query
func (s *FirstMileBigbag) QueryLazadaBigbagInfo (userInfo firstmilebigbagentity.QueryLazadaBigbagInfoUserInfoRequestEntity,client string,orderCode string,remark string,locale string,trackingNumber string) firstmilebigbagentity.QueryLazadaBigbagInfoResult {
    method := "/logistics/cnpms/bigbag/query"
    params := lib.InRow{
      "userInfo":userInfo,
      "client":client,
    }
    if orderCode!=""{
        params["orderCode"]=orderCode
    }
    if remark!=""{
        params["remark"]=remark
    }
    if locale!=""{
        params["locale"]=locale
    }
    if trackingNumber!=""{
        params["trackingNumber"]=trackingNumber
    }
    result := firstmilebigbagentity.QueryLazadaBigbagInfoResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//QueryPackageStatus
//@Title 1.seller post the tracking number and get the status of the package relative to the tracking number.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/package/query
func (s *FirstMileBigbag) QueryPackageStatus (trackingNumbers []string) firstmilebigbagentity.QueryPackageStatusResult {
    method := "/logistics/package/query"
    params := lib.InRow{
      "tracking_numbers":trackingNumbers,
    }
    result := firstmilebigbagentity.QueryPackageStatusResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UploadLazadaBagNumber
//@Title 1.seller call api to post the Lazada bag number and relative tracking number
//2.Api return the lzd bag number.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/lzd_bag/upload
func (s *FirstMileBigbag) UploadLazadaBagNumber (trackingNumbers []string,erpBagNumber string,parcelsTotal int,lzdBagNumber string) firstmilebigbagentity.UploadLazadaBagNumberResult {
    method := "/logistics/lzd_bag/upload"
    params := lib.InRow{
      "tracking_numbers":trackingNumbers,
      "erp_bag_number":erpBagNumber,
      "parcels_total":parcelsTotal,
      "lzd_bag_number":lzdBagNumber,
    }
    result := firstmilebigbagentity.UploadLazadaBagNumberResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UploadSellerBagNumber
//@Title 1.seller call api to post the seller bag number and relative tracking number
//2.Api return the lzd bag number.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/seller_bag/upload
func (s *FirstMileBigbag) UploadSellerBagNumber (trackingNumbers []string,erpBagNumber string,parcelsTotal int) firstmilebigbagentity.UploadSellerBagNumberResult {
    method := "/logistics/seller_bag/upload"
    params := lib.InRow{
      "tracking_numbers":trackingNumbers,
      "erp_bag_number":erpBagNumber,
      "parcels_total":parcelsTotal,
    }
    result := firstmilebigbagentity.UploadSellerBagNumberResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetChannelcodeByFirstMileNo
//@Title get channelcode by first mile No
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=34&path=/logistics/cngfc/fulfill/getchannelcode
func (s *FirstMileBigbag) GetChannelcodeByFirstMileNo (firstMileNos []string) firstmilebigbagentity.GetChannelcodeByFirstMileNoResult {
    method := "/logistics/cngfc/fulfill/getchannelcode"
    params := lib.InRow{
      "firstMileNos":firstMileNos,
    }
    result := firstmilebigbagentity.GetChannelcodeByFirstMileNoResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
