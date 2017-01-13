package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 门店插件上架操作
// 本接口用于上架商户门店订购的服务。
type AlipayOpenServicemarketCommodityShopOnlineRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                       `json:"terminal_type"`
  TerminalInfo      string                                                       `json:"terminal_info"`
  ProdCode          string                                                       `json:"prod_code"`
  NotifyUrl         string                                                       `json:"notify_url"`
  ReturnUrl         string                                                       `json:"return_url"`
  BizContent        AlipayOpenServicemarketCommodityShopOnlineRequestBizContent  `json:"biz_content"`
}

type AlipayOpenServicemarketCommodityShopOnlineRequestBizContent struct {

}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.commodity.shop.online"
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
