package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 活动下架接口
// 商户可以下架活动
type KoubeiMarketingCampaignActivityOfflineRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                   `json:"terminal_type"`
  TerminalInfo      string                                                   `json:"terminal_info"`
  ProdCode          string                                                   `json:"prod_code"`
  NotifyUrl         string                                                   `json:"notify_url"`
  ReturnUrl         string                                                   `json:"return_url"`
  BizContent        KoubeiMarketingCampaignActivityOfflineRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignActivityOfflineRequestBizContent struct {
  OutBizNo      string `json:"out_biz_no"`    // 外部批次ID,每次需传入不同的值
  CampId        string `json:"camp_id"`       // 活动Id
  Reason        string `json:"reason"`        // 下架原因
  ExtInfo       string `json:"ext_info"`      // 下架活动的扩展信息，不需要设置
  OperatorId    string `json:"operator_id"`   // 操作人id，与operator_type必须配对存在，当不填的时候默认是商户
  OperatorType  string `json:"operator_type"` // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.activity.offline"
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignActivityOfflineRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
