package coupon

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 卡券模板创建
// 创建卡券的模板，卡券的样式、内容信息通过该接口提交到支付宝，支付宝会生成模板ID反馈给开发者，用于后续的更新和发布。
type AlipayPassTemplateAddRequest struct {
  api.IAlipayRequest
  TerminalType      string                                   `json:"terminal_type"`
  TerminalInfo      string                                   `json:"terminal_info"`
  ProdCode          string                                   `json:"prod_code"`
  NotifyUrl         string                                   `json:"notify_url"`
  ReturnUrl         string                                   `json:"return_url"`
  BizContent        AlipayPassTemplateAddRequestBizContent   `json:"biz_content"`
}

type AlipayPassTemplateAddRequestBizContent struct {
  UniqueId    string `json:"unique_id"`     // 商户用于控制模版的唯一性。（可以使用时间戳保证唯一性）
  TplContent  string `json:"tpl_content"`   // 模板内容信息，遵循JSON规范，详情参见tpl_content参数说明(https://doc.open.alipay.com/doc2/detail.htm?treeId=193&articleId=105249&docType=1#tpl_content)
}

func (this *AlipayPassTemplateAddRequest) GetApiMethodName() string {
  return "alipay.pass.template.add"
}

func (this *AlipayPassTemplateAddRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayPassTemplateAddRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayPassTemplateAddRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayPassTemplateAddRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayPassTemplateAddRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayPassTemplateAddRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayPassTemplateAddRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayPassTemplateAddRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
