package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 营销活动人群组规则详情查询接口
// 口碑商户人群组详情查询接口
type KoubeiMarketingCampaignCrowdDetailQueryRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                    `json:"terminal_type"`
  TerminalInfo        string                                                    `json:"terminal_info"`
  ProdCode            string                                                    `json:"prod_code"`
  NotifyUrl           string                                                    `json:"notify_url"`
  ReturnUrl           string                                                    `json:"return_url"`
  BizContent          KoubeiMarketingCampaignCrowdDetailQueryRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignCrowdDetailQueryRequestBizContent struct {
  CrowdGroupId string `json:"crowd_group_id"` // 人群组ID，人群组创建成功时返回的ID
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.crowd.detail.query"
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignCrowdDetailQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
