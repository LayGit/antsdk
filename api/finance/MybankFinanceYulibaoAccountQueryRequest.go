package finance

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 余利宝账户和收益查询
// 余利宝账户（总份额、可用余额、冻结份额等）及收益（包括近1日、1周、历史累计收益等）查询
type MybankFinanceYulibaoAccountQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                              `json:"terminal_type"`
  TerminalInfo      string                                              `json:"terminal_info"`
  ProdCode          string                                              `json:"prod_code"`
  NotifyUrl         string                                              `json:"notify_url"`
  ReturnUrl         string                                              `json:"return_url"`
  BizContent        MybankFinanceYulibaoAccountQueryRequestBizContent   `json:"biz_content"`
}

type MybankFinanceYulibaoAccountQueryRequestBizContent struct {
  FundCode string `json:"fund_code"` // 基金代码，必填。目前默认填001529，代表余利宝
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetApiMethodName() string {
  return "mybank.finance.yulibao.account.query"
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *MybankFinanceYulibaoAccountQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *MybankFinanceYulibaoAccountQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
