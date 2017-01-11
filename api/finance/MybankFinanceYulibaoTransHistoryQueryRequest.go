package finance

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 余利宝历史交易查询
type MybankFinanceYulibaoTransHistoryQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                  `json:"terminal_type"`
  TerminalInfo      string                                                  `json:"terminal_info"`
  ProdCode          string                                                  `json:"prod_code"`
  NotifyUrl         string                                                  `json:"notify_url"`
  ReturnUrl         string                                                  `json:"return_url"`
  BizContent        MybankFinanceYulibaoTransHistoryQueryRequestBizContent  `json:"biz_content"`
}

type MybankFinanceYulibaoTransHistoryQueryRequestBizContent struct {
  FundCode  string  `json:"fund_code"`      // 基金代码，必填。目前默认填001529，代表余利宝。
  StartDate string  `json:"start_date"`     // 查询交易的开始时间，必须是格式为yyyyMMdd的日期字符串，如20160808。
  EndDate   string  `json:"end_date"`       // 查询交易的结束时间，必须是格式为yyyyMMdd的日期字符串，且日期要大于等于start_date，时间最大跨度为30天，如start_date为20160808，则end_date最大值为20160906。
  Page      int     `json:"page"`           // 页码，历史交易记录分页展示的页码。必须为正整数，最大值为99。
  PageSize  int     `json:"page_size"`      // 每页条数，历史交易记录查询时每页的最大条数。必须为正整数，最大值为50。
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetApiMethodName() string {
  return "mybank.finance.yulibao.trans.history.query"
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *MybankFinanceYulibaoTransHistoryQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
