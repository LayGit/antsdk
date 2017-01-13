package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 订购插件订单明细查询
// 第三方服务商提供服务产品被商户订购后，服务市场会推送订单信息给服务商，服务商根据订单号回查该订单明细信息。
type AlipayOpenServicemarketOrderQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                               `json:"terminal_type"`
  TerminalInfo      string                                               `json:"terminal_info"`
  ProdCode          string                                               `json:"prod_code"`
  NotifyUrl         string                                               `json:"notify_url"`
  ReturnUrl         string                                               `json:"return_url"`
  BizContent        AlipayOpenServicemarketOrderQueryRequestBizContent   `json:"biz_content"`
}

type AlipayOpenServicemarketOrderQueryRequestBizContent struct {
  CommodityOrderId  string `json:"commodity_order_id"`    // 商户订单ID号
  StartPage         string `json:"start_page"`            // 从第几页开始查询
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.order.query"
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketOrderQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
