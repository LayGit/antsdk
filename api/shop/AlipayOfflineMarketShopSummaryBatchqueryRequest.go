package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 门店摘要信息批量查询接口
// 用于进行门店摘要信息批量查询接口。
type AlipayOfflineMarketShopSummaryBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                                                      `json:"terminal_type"`
  TerminalInfo  string                                                      `json:"terminal_info"`
  ProdCode      string                                                      `json:"prod_code"`
  NotifyUrl     string                                                      `json:"notify_url"`
  ReturnUrl     string                                                      `json:"return_url"`
  BizContent    AlipayOfflineMarketShopSummaryBatchqueryRequestBizContent   `json:"biz_content"`
}

type AlipayOfflineMarketShopSummaryBatchqueryRequestBizContent struct {
  OPRole            string  `json:"op_role"`              // 表示接口业务的调用方身份：ISV、 服务商身份标识。传入ISV代表系统集成商身份。传入PROVIDER代表服务商。
  QueryType         string  `json:"query_type"`           // 门店数据查询类型，根据类型可以返回指定的门店数据，目前支持的类型如下： BRAND_RELATION ： 品牌商关联店铺 MALL_SELF ：MALL自己的门店 MALL_RELATION：MALL关联下的门店 MERCHANT_SELF:商户自己的门店
  RelatedPartnerId  string  `json:"related_partner_id"`   // query_type查询类型下所关联的商户PID
  ShopId            string  `json:"shop_id"`              // 门店ID
  ShopStatus        string  `json:"shop_status"`          // 门店状态，传入多个状态，多个状态使用英文逗号隔开，例如：PAUSED,OPEN 店铺状态：OPEN（营业）、PAUSED（暂停）、INIT（初始）、FREEZE（冻结）、CLOSED（关店）
  PageNo            int     `json:"page_no"`              // 页码，留空标示第一页，默认 20个结果为一页
  PageSize          int     `json:"page_size"`            // 每页记录数，默认20，最大 100
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetApiMethodName() string {
  return "alipay.offline.market.shop.summary.batchquery"
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMarketShopSummaryBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
