package bill

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

type AlipayDataDataserviceBillDownloadurlQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                        `json:"terminal_type"`
  TerminalInfo      string                                                        `json:"terminal_info"`
  ProdCode          string                                                        `json:"prod_code"`
  NotifyUrl         string                                                        `json:"notify_url"`
  ReturnUrl         string                                                        `json:"return_url"`
  BizContent        AlipayDataDataserviceBillDownloadurlQueryRequestBizContent    `json:"biz_content"`
}

type AlipayDataDataserviceBillDownloadurlQueryRequestBizContent struct {
  BillType string `json:"bill_type"` // 账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型：trade、signcustomer；trade指商户基于支付宝交易收单的业务账单；signcustomer是指基于商户支付宝余额收入及支出等资金变动的帐务账单；
  BillDate string `json:"bill_date"` // 账单时间：日账单格式为yyyy-MM-dd，月账单格式为yyyy-MM。
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetApiMethodName() string {
  return "alipay.data.dataservice.bill.downloadurl.query"
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayDataDataserviceBillDownloadurlQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
