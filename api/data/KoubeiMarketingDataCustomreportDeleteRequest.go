package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 自定义报表规则删除接口
// 删除自定义数据报表规则
type KoubeiMarketingDataCustomreportDeleteRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                  `json:"terminal_type"`
  TerminalInfo        string                                                  `json:"terminal_info"`
  ProdCode            string                                                  `json:"prod_code"`
  NotifyUrl           string                                                  `json:"notify_url"`
  ReturnUrl           string                                                  `json:"return_url"`
  BizContent          KoubeiMarketingDataCustomreportDeleteRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportDeleteRequestBizContent struct {
  ConditionKey string `json:"condition_key"` // 自定义报表规则的KEY
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetApiMethodName() string {
  return "koubei.marketing.data.customreport.delete"
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
