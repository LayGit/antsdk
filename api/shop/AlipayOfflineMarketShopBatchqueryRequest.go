package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 查询商户的门店编号列表
// 系统商通过该接口可以查询对应APPID下所有的门店编号（支付宝口碑门店编号）
type AlipayOfflineMarketShopBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                                              `json:"terminal_type"`
  TerminalInfo  string                                              `json:"terminal_info"`
  ProdCode      string                                              `json:"prod_code"`
  NotifyUrl     string                                              `json:"notify_url"`
  ReturnUrl     string                                              `json:"return_url"`
  BizContent    AlipayOfflineMarketShopBatchqueryRequestBizContent  `json:"biz_content"`
}

type AlipayOfflineMarketShopBatchqueryRequestBizContent struct {
  PageNo  string `json:"page_no"` // 页码，第一页传入"1"，默认500个结果为一页
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetApiMethodName() string {
  return "alipay.offline.market.shop.batchquery"
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
