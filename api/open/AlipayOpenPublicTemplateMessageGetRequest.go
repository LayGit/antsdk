package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 消息模板领取接口
// 帮助服务窗开发者从服务窗平台（fuwu.alipay.com）公共模板库里领取一个行业的消息模板--消息模板是一种消息的样式，如消费提醒、账单提醒等
type AlipayOpenPublicTemplateMessageGetRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                `json:"terminal_type"`
  TerminalInfo      string                                                `json:"terminal_info"`
  ProdCode          string                                                `json:"prod_code"`
  NotifyUrl         string                                                `json:"notify_url"`
  ReturnUrl         string                                                `json:"return_url"`
  BizContent        AlipayOpenPublicTemplateMessageGetRequestBizContent   `json:"biz_content"`
}

type AlipayOpenPublicTemplateMessageGetRequestBizContent struct {
  TemplateId string `json:"template_id"`  // 消息母板id，登陆生活号后台(fuwu.alipay.com)，点击菜单“模板消息”，点击“模板库”，即可看到相应模板的消息母板id
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetApiMethodName() string {
  return "alipay.open.public.template.message.get"
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicTemplateMessageGetRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
