package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 自定义报表数据查询接口
// 查询自定义数据报表数据
type KoubeiMarketingDataCustomreportQueryRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                 `json:"terminal_type"`
  TerminalInfo        string                                                 `json:"terminal_info"`
  ProdCode            string                                                 `json:"prod_code"`
  NotifyUrl           string                                                 `json:"notify_url"`
  ReturnUrl           string                                                 `json:"return_url"`
  BizContent          KoubeiMarketingDataCustomreportQueryRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportQueryRequestBizContent struct {
  ConditionKey  string      `json:"condition_key"`  // 规则KEY
  FilterTags    []FilterTag `json:"filter_tags"`    // 额外增加的查询过滤条件
  MaxCount      string      `json:"max_count"`      // 一次拉多少条
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetApiMethodName() string {
  return "koubei.marketing.data.customreport.query"
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingDataCustomreportQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
