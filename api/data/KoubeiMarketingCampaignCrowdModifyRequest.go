package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 营销活动人群组规则修改接口
// 口碑商户人群组修改接口
type KoubeiMarketingCampaignCrowdModifyRequest struct {
  api.IAlipayRequest
  TerminalType        string                                               `json:"terminal_type"`
  TerminalInfo        string                                               `json:"terminal_info"`
  ProdCode            string                                               `json:"prod_code"`
  NotifyUrl           string                                               `json:"notify_url"`
  ReturnUrl           string                                               `json:"return_url"`
  BizContent          KoubeiMarketingCampaignCrowdModifyRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignCrowdModifyRequestBizContent struct {
  OutBizNo      string `json:"out_biz_no"`      // 外部流水号
  CrowdGroupId  string `json:"crowd_group_id"`  // 人群组的唯一标识ID
  Conditions    string `json:"conditions"`      // 圈人的条件 op:表示操作符，目前支持EQ相等,GT大于,GTEQ大于等于,LT小于,LTEQ小于等于,NEQ不等,LIKE模糊匹配,IN在枚举范围内,NOTIN不在枚举范围内,BETWEEN范围比较,LEFTDAYS几天以内,RIGHTDAYS几天以外,LOCATE地理位置比较,LBS地图位置数据 tagCode:标签code，详细标签code参见附件。标签信息 value:标签值
  OperatorId    string `json:"operator_id"`     // 操作人id，必须和operator_type配对出现，不填时默认是商户
  OperatorType  string `json:"operator_type"`   // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.crowd.modify"
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignCrowdModifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
