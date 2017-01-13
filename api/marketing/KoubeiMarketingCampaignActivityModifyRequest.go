package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 活动修改接口
// 活动修改，可修改一下几点： 活动结束时间/券有效期 —— 只可延长 活动库存 —— 只可追加 活动参与限制（包含每月/周/日) —— 只可追加 活动门店/券适用门店 —— 只可追加
type KoubeiMarketingCampaignActivityModifyRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                  `json:"terminal_type"`
  TerminalInfo      string                                                  `json:"terminal_info"`
  ProdCode          string                                                  `json:"prod_code"`
  NotifyUrl         string                                                  `json:"notify_url"`
  ReturnUrl         string                                                  `json:"return_url"`
  BizContent        KoubeiMarketingCampaignActivityModifyRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignActivityModifyRequestBizContent struct {
  OutBizNo        string            `json:"out_biz_no"`       // 外部批次ID，用户指定,每次请求保持唯一
  OperatorId      string            `json:"operator_id"`      // 操作人id，必须和operator_type配对出现，不填时默认是商户
  OperatorType    string            `json:"operator_type"`    // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
  CampId          string            `json:"camp_id"`          // 活动id
  Name            string            `json:"name"`             // 活动名称 不允许修改，必须与活动详情查询的结果保持一致
  StartTime       string            `json:"start_time"`       // 活动开始时间 不允许修改，必须与活动详情查询的结果保持一致
  EndTime         string            `json:"end_time"`         // 活动结束时间 活动结束时间只允许延长
  Type            string            `json:"type"`             // 活动类型 不允许修改，必须与活动详情查询的结果保持一致
  Desc            string            `json:"desc"`             // 活动详细说明 不允许修改，必须与活动详情查询的结果保持一致
  BudgetInfo      BudgetInfo        `json:"budget_info"`      // 活动预算
  AutoDelayFlag   string            `json:"auto_delay_flag"`  // 是否自动续期活动，仅当活动下券的有效期均为相对有效期时才能设置成Y
  ConstraintInfo  ConstraintInfo    `json:"constraint_info"`  // 活动限制信息
  PromoTools      []PromoTool       `json:"promo_tools"`      // 营销工具集
  PublishChannels []PublishChannel  `json:"publish_channels"` // 投放渠道集，当活动类型为DIRECT_SEND或者REAL_TIME_SEND时必填，为CONSUME_SEND时必须为空
  RecruitTool     RecruitTool       `json:"recruit_tool"`     // 招商工具
  ExtInfo         string            `json:"ext_info"`         // 活动的扩展信息
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.activity.modify"
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignActivityModifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
