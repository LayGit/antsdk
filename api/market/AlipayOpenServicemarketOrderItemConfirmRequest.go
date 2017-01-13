package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务商代商家确认实施完成
// 此接口需商家授权服务商应用权限后，服务商可代商家进行实施确认完成动作
type AlipayOpenServicemarketOrderItemConfirmRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                    `json:"terminal_type"`
  TerminalInfo      string                                                    `json:"terminal_info"`
  ProdCode          string                                                    `json:"prod_code"`
  NotifyUrl         string                                                    `json:"notify_url"`
  ReturnUrl         string                                                    `json:"return_url"`
  BizContent        AlipayOpenServicemarketOrderItemConfirmRequestBizContent  `json:"biz_content"`
}

type AlipayOpenServicemarketOrderItemConfirmRequestBizContent struct {
  CommodityOrderId  string `json:"commodity_order_id"`  // 商品订单ID
  ShopId            string `json:"shop_id"`             // 商家订购服务选择的某一门店的ID
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.order.item.confirm"
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketOrderItemConfirmRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
