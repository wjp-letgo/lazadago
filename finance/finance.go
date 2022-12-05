package finance

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	financeentity "github.com/wjpxxx/lazadago/finance/entity"
)
//Finance
type Finance struct{
	Config *lazadaConfig.Config
}

//GetPayoutStatus
//@Title Get your transaction statements  created after the provided date
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=9&path=/finance/payout/status/get
func (s *Finance) GetPayoutStatus (createdAfter string) financeentity.GetPayoutStatusResult {
    method := "/finance/payout/status/get"
    params := lib.InRow{
      "created_after":createdAfter,
    }
    result := financeentity.GetPayoutStatusResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetTransactionDetails
//@Title Use this API to get transaction or fee details in a specified period.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=9&path=/finance/transaction/detail/get
func (s *Finance) GetTransactionDetails (transType string,startTime string,endTime string,limit string,offset string) financeentity.GetTransactionDetailsResult {
    method := "/finance/transaction/detail/get"
    params := lib.InRow{
      "start_time":startTime,
      "end_time":endTime,
    }
    if transType!=""{
        params["trans_type"]=transType
    }
    if limit!=""{
        params["limit"]=limit
    }
    if offset!=""{
        params["offset"]=offset
    }
    result := financeentity.GetTransactionDetailsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//QueryTransactionDetails
//@Title API to query seller transaction details within specific date range.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=9&path=/finance/transaction/details/get
func (s *Finance) QueryTransactionDetails (offset string,transType string,tradeOrderId string,limit string,startTime string,endTime string,tradeOrderLineId string) financeentity.QueryTransactionDetailsResult {
    method := "/finance/transaction/details/get"
    params := lib.InRow{
      "start_time":startTime,
      "end_time":endTime,
    }
    if offset!=""{
        params["offset"]=offset
    }
    if transType!=""{
        params["trans_type"]=transType
    }
    if tradeOrderId!=""{
        params["trade_order_id"]=tradeOrderId
    }
    if limit!=""{
        params["limit"]=limit
    }
    if tradeOrderLineId!=""{
        params["trade_order_line_id"]=tradeOrderLineId
    }
    result := financeentity.QueryTransactionDetailsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
