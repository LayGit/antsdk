package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务订单明细实施项单项取消
// 当服务商订购服务区域较多时，可能部分区域不能实施，可以取消实施
type AlipayOpenServicemarketOrderItemCancelRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                  `json:"terminal_type"`
  TerminalInfo      string                                                  `json:"terminal_info"`
  ProdCode          string                                                  `json:"prod_code"`
  NotifyUrl         string                                                  `json:"notify_url"`
  ReturnUrl         string                                                  `json:"return_url"`
  BizContent        AlipayOpenServicemarketOrderItemCancelRequestBizContent `json:"biz_content"`
}

type AlipayOpenServicemarketOrderItemCancelRequestBizContent struct {
  CommodityOrderId  string `json:"commodity_order_id"`  // 订购服务订单ID
  ShopId            string `json:"shop_id"`             // 订购服务门店ID
  CancelReason      string `json:"cancel_reason"`       // 当前门店区域不支持实施
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.order.item.cancel"
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
