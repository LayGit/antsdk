package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 营销活动人群组规则删除接口
// 口碑商户人群组删除接口
type KoubeiMarketingCampaignCrowdDeleteRequest struct {
  api.IAlipayRequest
  TerminalType        string                                               `json:"terminal_type"`
  TerminalInfo        string                                               `json:"terminal_info"`
  ProdCode            string                                               `json:"prod_code"`
  NotifyUrl           string                                               `json:"notify_url"`
  ReturnUrl           string                                               `json:"return_url"`
  BizContent          KoubeiMarketingCampaignCrowdDeleteRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignCrowdDeleteRequestBizContent struct {
  CrowdGroupId string `json:"crowd_group_id"` // 人群组的唯一标识ID
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.crowd.delete"
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignCrowdDeleteRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
