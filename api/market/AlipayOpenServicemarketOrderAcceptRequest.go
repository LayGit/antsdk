package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务商接单操作
// 对于商户订购的服务插件，插件服务商需要接单之后，才能有后续的操作
type AlipayOpenServicemarketOrderAcceptRequest struct {
  api.IAlipayRequest
  TerminalType      string                                               `json:"terminal_type"`
  TerminalInfo      string                                               `json:"terminal_info"`
  ProdCode          string                                               `json:"prod_code"`
  NotifyUrl         string                                               `json:"notify_url"`
  ReturnUrl         string                                               `json:"return_url"`
  BizContent        AlipayOpenServicemarketOrderQueryRequestBizContent   `json:"biz_content"`
}

type AlipayOpenServicemarketOrderAcceptRequestBizContent struct {
  CommodityOrderId string `json:"commodity_order_id"` // 服务商品订单ID
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.order.accept"
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketOrderAcceptRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
