package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 会员卡模板查询接口
type AlipayMarketingCardTemplateQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                             `json:"terminal_type"`
  TerminalInfo      string                                             `json:"terminal_info"`
  ProdCode          string                                             `json:"prod_code"`
  NotifyUrl         string                                             `json:"notify_url"`
  ReturnUrl         string                                             `json:"return_url"`
  BizContent        AlipayMarketingCardTemplateQueryRequestBizContent  `json:"biz_content"`
}

type AlipayMarketingCardTemplateQueryRequestBizContent struct {
  TemplateId string `json:"template_id"`  // 支付宝卡模板ID（模板创建接口返回的支付宝端模板ID）
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetApiMethodName() string {
  return "alipay.marketing.card.template.query"
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCardTemplateQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
