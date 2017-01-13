package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 口碑营销活动列表查询
// 查询活动列表
type KoubeiMarketingCampaignActivityBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                      `json:"terminal_type"`
  TerminalInfo      string                                                      `json:"terminal_info"`
  ProdCode          string                                                      `json:"prod_code"`
  NotifyUrl         string                                                      `json:"notify_url"`
  ReturnUrl         string                                                      `json:"return_url"`
  BizContent        KoubeiMarketingCampaignActivityBatchqueryRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingCampaignActivityBatchqueryRequestBizContent struct {
  QueryCriterias  []Condition `json:"query_criterias"`  // 查询条件
  PageNumber      string      `json:"page_number"`      // 页码，默认为1
  OperatorId      string      `json:"operator_id"`      // 操作人id，必须和operator_type配对存在，不填时默认是商户
  OperatorType    string      `json:"operator_type"`    // 操作人类型,有以下值可填：MER（外部商户），MER_OPERATOR（外部商户操作员），PROVIDER（外部服务商），PROVIDER_STAFF（外部服务商员工），默认不需要填这个字段，默认为MER
  PageSize        string      `json:"page_size"`        // 页大小，默认为20
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.activity.batchquery"
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignActivityBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
