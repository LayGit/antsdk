package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 营销活动人群组人数统计接口
// 口碑商户人群组数目统计接口
type KoubeiMarketingCampaignCrowdCountRequest struct {
  api.IAlipayRequest
  TerminalType        string                                              `json:"terminal_type"`
  TerminalInfo        string                                              `json:"terminal_info"`
  ProdCode            string                                              `json:"prod_code"`
  NotifyUrl           string                                              `json:"notify_url"`
  ReturnUrl           string                                              `json:"return_url"`
  BizContent          KoubeiMarketingCampaignCrowdCountRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignCrowdCountRequestBizContent struct {
  CrowdGroupId  string `json:"crowd_group_id"`  // crowd_group_id和conditions不能同时为空 如果crowd_group_id不为空则根据crowd_group_id查询人群分组的信息进行统计，否则以conditions的内容为过滤条件进行统计，如果crowd_group_id和conditions都不为空则优先使用conditions的条件
  Conditions    string `json:"conditions"`      // 圈人的条件 op:表示操作符，目前支持EQ相等,GT大于,GTEQ大于等于,LT小于,LTEQ小于等于,NEQ不等,LIKE模糊匹配,IN在枚举范围内,NOTIN不在枚举范围内,BETWEEN范围比较,LEFTDAYS几天以内,RIGHTDAYS几天以外,LOCATE地理位置比较,LBS地图位置数据 tagCode:标签code，详细标签code参见附件。标签信息 value:标签值
  Dimensions    string `json:"dimensions"`      // 画像分析的维度，目前支持:["pam_age","pam_gender","pam_constellation","pam_hometown_code","pam_city_code","pam_occupation","pam_consume_level","pam_have_baby"]，以koubei.marketing.campaign.tags.query接口返回的dimensions为准，各个维度标签的详细信息参见附件，标签信息
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.crowd.count"
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignCrowdCountRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
