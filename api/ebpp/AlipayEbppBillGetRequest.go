package ebpp

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 查询账单
type AlipayEbppBillGetRequest struct {
  api.IAlipayRequest
  TerminalType      string  `json:"terminal_type"`
  TerminalInfo      string  `json:"terminal_info"`
  ProdCode          string  `json:"prod_code"`
  NotifyUrl         string  `json:"notify_url"`
  ReturnUrl         string  `json:"return_url"`
  OrderType         string  `json:"order_type"`
  MerchantOrderNo   string  `json:"merchant_order_no"`
}

func (this *AlipayEbppBillGetRequest) GetApiMethodName() string {
  return "alipay.ebpp.bill.get"
}

func (this *AlipayEbppBillGetRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEbppBillGetRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEbppBillGetRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEbppBillGetRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEbppBillGetRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEbppBillGetRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEbppBillGetRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEbppBillGetRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("order_type", this.OrderType)
  txtParams.Put("merchant_order_no", this.MerchantOrderNo)
  return txtParams
}
