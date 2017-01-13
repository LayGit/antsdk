package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 营销活动人群组规则创建接口
// 口碑商户人群组创建接口
type KoubeiMarketingCampaignCrowdCreateRequest struct {
  api.IAlipayRequest
  TerminalType        string                                               `json:"terminal_type"`
  TerminalInfo        string                                               `json:"terminal_info"`
  ProdCode            string                                               `json:"prod_code"`
  NotifyUrl           string                                               `json:"notify_url"`
  ReturnUrl           string                                               `json:"return_url"`
  BizContent          KoubeiMarketingCampaignCrowdCreateRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignCrowdCreateRequestBizContent struct {
  OutBizNo      string `json:"out_biz_no"`    // 外部流水号
  Name          string `json:"name"`          // 人群组的名称
  Conditions    string `json:"conditions"`    // 圈人的条件 op:表示操作符，目前支持EQ相等,GT大于,GTEQ大于等于,LT小于,LTEQ小于等于,NEQ不等,LIKE模糊匹配,IN在枚举范围内,NOTIN不在枚举范围内,BETWEEN范围比较,LEFTDAYS几天以内,RIGHTDAYS几天以外,LOCATE地理位置比较,LBS地图位置数据 tagCode:标签code，详细标签code参见附件。标签信息 value:标签值
  OperatorId    string `json:"operator_id"`   // 操作人id，必须和operator_type配对出现，不填时默认是商户
  OperatorType  string `json:"operator_type"` // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.crowd.create"
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignCrowdCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
