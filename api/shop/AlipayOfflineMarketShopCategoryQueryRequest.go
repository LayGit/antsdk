package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 门店类目配置查询接口
// 用于查询可用于开店的类目，以及类目上的配置约束信息
type AlipayOfflineMarketShopCategoryQueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                                                `json:"terminal_type"`
  TerminalInfo  string                                                `json:"terminal_info"`
  ProdCode      string                                                `json:"prod_code"`
  NotifyUrl     string                                                `json:"notify_url"`
  ReturnUrl     string                                                `json:"return_url"`
  BizContent    AlipayOfflineMarketShopCategoryQueryRequestBizContent `json:"biz_content"`
}

type AlipayOfflineMarketShopCategoryQueryRequestBizContent struct {
  CategoryId  string `json:"category_id"` // 类目ID，如果为空则查询全部类目。
  OPRole      string `json:"op_role"`     // 表示接口业务的调用方身份,默认不填标识为ISV。
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetApiMethodName() string {
  return "alipay.offline.market.shop.category.query"
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
