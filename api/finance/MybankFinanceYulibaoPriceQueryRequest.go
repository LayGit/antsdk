package finance

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

type MybankFinanceYulibaoPriceQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                            `json:"terminal_type"`
  TerminalInfo      string                                            `json:"terminal_info"`
  ProdCode          string                                            `json:"prod_code"`
  NotifyUrl         string                                            `json:"notify_url"`
  ReturnUrl         string                                            `json:"return_url"`
  BizContent        MybankFinanceYulibaoPriceQueryRequestBizContent   `json:"biz_content"`
}

type MybankFinanceYulibaoPriceQueryRequestBizContent struct {
  FundCode  string `json:"fund_code"`   // 基金代码，必填。目前默认填001529，代表余利宝。
  StartDate string `json:"start_date"`  // 查询行情的开始日期，必须是格式为yyyyMMdd的日期字符串，如20160808
  EndDate   string `json:"end_date"`    // 查询行情的截止日期，必须是格式为yyyyMMdd的日期字符串且日期要大于等于start_date，时间最大跨度为30天，如start_date为20160808，则end_date最大值为20160906
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetApiMethodName() string {
  return "mybank.finance.yulibao.price.query"
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *MybankFinanceYulibaoPriceQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *MybankFinanceYulibaoPriceQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
