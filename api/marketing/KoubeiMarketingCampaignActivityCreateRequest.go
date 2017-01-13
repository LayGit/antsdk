package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 活动创建接口
// 商户可以创建运营活动
type KoubeiMarketingCampaignActivityCreateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                  `json:"terminal_type"`
  TerminalInfo      string                                                  `json:"terminal_info"`
  ProdCode          string                                                  `json:"prod_code"`
  NotifyUrl         string                                                  `json:"notify_url"`
  ReturnUrl         string                                                  `json:"return_url"`
  BizContent        KoubeiMarketingCampaignActivityCreateRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignActivityCreateRequestBizContent struct {
  OutBizNo        string            `json:"out_biz_no"`         // 外部批次ID，同一owner创建活动需要换out_biz_no，控制幂等
  OperatorId      string            `json:"operator_id"`        // 操作人id，必须和operator_type配对出现，不填时默认是商户
  OperatorType    string            `json:"operator_type"`      // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
  Name            string            `json:"name"`               // 活动名称
  StartTime       string            `json:"start_time"`         // 活动开始时间
  EndTime         string            `json:"end_time"`           // 活动结束时间
  Type            string            `json:"type"`               // 活动类型，目前支持以下类型： CONSUME_SEND：消费送活动 DIRECT_SEND：直发奖活动 REAL_TIME_SEND：实时立减类活动 GUESS_SEND：口令送 RECHARGE_SEND：充值送 POINT_SEND：集点卡活动
  Desc            string            `json:"desc"`               // 活动详细说明
  BudgetInfo      BudgetInfo        `json:"budget_info"`        // 活动预算（当活动类型为GUESS_SEND且营销工具PromoTool的数量大于1时，即口令随机送活动，活动预算为空，每张券的预算数量落在SendRule里的send_budget）
  ConstraintInfo  ConstraintInfo    `json:"constraint_info"`    // 活动限制信息
  AutoDelayFlag   string            `json:"auto_delay_flag"`    // 是否自动续期活动，默认为N,只有当对应营销工具券有效期为相对有效期时才能设置成Y
  PromoTools      []PromoTool       `json:"promo_tools"`        // 营销工具集
  PublishChannels []PublishChannel  `json:"publish_channels"`   // 投放渠道集，当活动类型为DIRECT_SEND或者REAL_TIME_SEND时必填，当活动类型为CONSUME_SEND时必须为空，当活动类型为GUESS_SEND时，投放渠道只能有一个且type只能为MERCHANT_CROWD
  RecruitTool     RecruitTool       `json:"recruit_tool"`       // 招商工具信息
  ExtInfo         string            `json:"ext_info"`           // 活动的扩展信息，无需设置
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.activity.create"
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignActivityCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
