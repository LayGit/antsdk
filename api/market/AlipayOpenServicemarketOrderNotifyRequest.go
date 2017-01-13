package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务市场商户确认订购通知
// 服务市场当商户选择服务商提供产品并订购确认时，通知服务商订单消息。服务商可以通过通知的消息内容回查该订单明细。回查接口（alipay.open.servicemarket.order.query）
type AlipayOpenServicemarketOrderNotifyRequest struct {
  api.IAlipayRequest
  TerminalType      string                                               `json:"terminal_type"`
  TerminalInfo      string                                               `json:"terminal_info"`
  ProdCode          string                                               `json:"prod_code"`
  NotifyUrl         string                                               `json:"notify_url"`
  ReturnUrl         string                                               `json:"return_url"`
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.order.notify"
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
