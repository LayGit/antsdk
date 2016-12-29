package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 查询单个门店信息接口
// 系统商通过该接口可以查询单个门店的详细信息，包括门店基础信息，门店状态等信息
type AlipayOfflineMarketShopQuerydetailRequest struct {
  api.IAlipayRequest
  TerminalType  string                                              `json:"terminal_type"`
  TerminalInfo  string                                              `json:"terminal_info"`
  ProdCode      string                                              `json:"prod_code"`
  NotifyUrl     string                                              `json:"notify_url"`
  ReturnUrl     string                                              `json:"return_url"`
  BizContent    AlipayOfflineMarketShopQuerydetailRequestBizContent `json:"biz_content"`
}

type AlipayOfflineMarketShopQuerydetailRequestBizContent struct {
  ShopId string `json:"shop_id"`  // 支付宝门店ID
  OPRole string `json:"op_role"`  // 服务商及商户调用情况下务必传递。操作人角色，默认商户操作:MERCHANT；服务商操作：PROVIDER；ISV: 不需要填写
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetApiMethodName() string {
  return "alipay.offline.market.shop.querydetail"
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMarketShopQuerydetailRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
