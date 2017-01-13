package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 活动详情查询
type KoubeiMarketingCampaignActivityQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                 `json:"terminal_type"`
  TerminalInfo      string                                                 `json:"terminal_info"`
  ProdCode          string                                                 `json:"prod_code"`
  NotifyUrl         string                                                 `json:"notify_url"`
  ReturnUrl         string                                                 `json:"return_url"`
  BizContent        KoubeiMarketingCampaignActivityQueryRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignActivityQueryRequestBizContent struct {
  CampId        string `json:"camp_id"`       // 活动id
  OperatorId    string `json:"operator_id"`   // 操作人id，必须和operator_type配对出现，不填时默认是商户
  OperatorType  string `json:"operator_type"` // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.activity.query"
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignActivityQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
