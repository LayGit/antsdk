package market

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 门店插件下架操作
// 本接口需要商户授权服务商门店操作权限后，服务商可对商户门店上的插件进行下架操作
type AlipayOpenServicemarketCommodityShopOfflineRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                        `json:"terminal_type"`
  TerminalInfo      string                                                        `json:"terminal_info"`
  ProdCode          string                                                        `json:"prod_code"`
  NotifyUrl         string                                                        `json:"notify_url"`
  ReturnUrl         string                                                        `json:"return_url"`
  BizContent        AlipayOpenServicemarketCommodityShopOfflineRequestBizContent  `json:"biz_content"`
}

type AlipayOpenServicemarketCommodityShopOfflineRequestBizContent struct {
  CommodityId string `json:"commodity_id"`  // 服务商户ID
  ShopId      string `json:"shop_id"`       // 门店ID
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetApiMethodName() string {
  return "alipay.open.servicemarket.commodity.shop.offline"
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenServicemarketCommodityShopOfflineRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
