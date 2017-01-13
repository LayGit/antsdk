package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 自定义报表规则详情查询接口
// 查询自定义数据报表规则详情
type KoubeiMarketingDataCustomreportDetailQueryRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                      `json:"terminal_type"`
  TerminalInfo        string                                                      `json:"terminal_info"`
  ProdCode            string                                                      `json:"prod_code"`
  NotifyUrl           string                                                      `json:"notify_url"`
  ReturnUrl           string                                                      `json:"return_url"`
  BizContent          KoubeiMarketingDataCustomreportDetailQueryRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportDetailQueryRequestBizContent struct {
  ConditionKey string `json:"condition_key"` // 自定义报表的规则KEY
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetApiMethodName() string {
  return "koubei.marketing.data.customreport.detail.query"
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
