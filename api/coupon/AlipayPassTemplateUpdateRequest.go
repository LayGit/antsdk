package coupon

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 支付宝pass更新模版接口
// 对于已经创建的模板，如果需要修改模板内容，可通过该接口修改，适用于修改模板内容；对于已经发布的模板，如果需要修改内容并同步到用户端，则应使用更新卡券接口。
type AlipayPassTemplateUpdateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                      `json:"terminal_type"`
  TerminalInfo      string                                      `json:"terminal_info"`
  ProdCode          string                                      `json:"prod_code"`
  NotifyUrl         string                                      `json:"notify_url"`
  ReturnUrl         string                                      `json:"return_url"`
  BizContent        AlipayPassTemplateUpdateRequestBizContent   `json:"biz_content"`
}

type AlipayPassTemplateUpdateRequestBizContent struct {
  TplId       string `json:"tpl_id"`        // 更新的模板ID
  TplContent  string `json:"tpl_content"`   // 模板内容信息，遵循JSON规范，详情参见tpl_content参数说明(https://doc.open.alipay.com/doc2/detail.htm?treeId=193&articleId=105249&docType=1#tpl_content)
}

func (this *AlipayPassTemplateUpdateRequest) GetApiMethodName() string {
  return "alipay.pass.template.update"
}

func (this *AlipayPassTemplateUpdateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayPassTemplateUpdateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayPassTemplateUpdateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayPassTemplateUpdateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayPassTemplateUpdateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayPassTemplateUpdateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayPassTemplateUpdateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayPassTemplateUpdateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
