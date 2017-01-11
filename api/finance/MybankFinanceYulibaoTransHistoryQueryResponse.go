package finance

import (
  "github.com/LayGit/antsdk/api"
)

type MybankFinanceYulibaoTransHistoryQueryResponse struct {
  api.AlipayResponse
  HistoryTransDetailInfos []YLBTransDetailInfo  `json:"history_trans_detail_infos"` // 历史交易详情信息
  CurrentPage             int                   `json:"current_page"`               // 历史交易记录查询的当前页码
  HasNextPage             bool                  `json:"has_next_page"`              // 当前查询是否具有下一页的数据，true-有，fasle-没有
  TotalItemCount          int                   `json:"total_item_count,string"`    // 当前查询在不分页情况下的数据总数
}

type YLBTransDetailInfo struct {
  Amount            float64 `json:"amount,string"`        // 余利宝交易金额，单位为元
  TransDate         string  `json:"trans_date"`           // 交易时间
  TransAccountDate  string  `json:"trans_account_date"`   // 交易到账时间
  BizNo             string  `json:"biz_no"`               // 交易流水号
  TransName         string  `json:"trans_name"`           // 交易名称，如正常申购、正常赎回、收益发放、份额强增、份额强减
  TransType         string  `json:"trans_type"`           // 交易类型，如正常申购、正常赎回、收益发放、份额强增、份额强减
  TransStatus       string  `json:"trans_status"`         // 交易状态，如success-成功、failure-失败、inprocess-进行中等
  Memo              string  `json:"memo"`                 // 备注信息
}
