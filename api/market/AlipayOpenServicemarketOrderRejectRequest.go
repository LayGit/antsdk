package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务商拒绝接单
// 商户在一个不能提供服务的区域订购了门店，当服务商所提供的服务不能触达的时候，服务商可以拒绝接单。
type AlipayOpenServicemarketOrderRejectRequest struct {
  api.IAlipayRequest
  TerminalType      string                                               `json:"terminal_type"`
  TerminalInfo      string                                               `json:"terminal_info"`
  ProdCode          string                                               `json:"prod_code"`
  NotifyUrl         string                                               `json:"notify_url"`
  ReturnUrl         string                                               `json:"return_url"`
  BizContent        AlipayOpenServicemarketOrderRejectRequestBizContent  `json:"biz_content"`
}

type AlipayOpenServicemarketOrderRejectRequestBizContent struct {
  CommodityOrderId  string `json:"commodity_order_id"`  // 订购服务商品订单ID
  RejectReason      string `json:"reject_reason"`       // v
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.order.reject"
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderRejectRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketOrderRejectRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
