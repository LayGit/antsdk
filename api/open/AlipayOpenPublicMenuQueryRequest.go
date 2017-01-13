package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 公众服务平台查询菜单
// 开发者通过调用该接口查询菜单。
type AlipayOpenPublicMenuQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                          `json:"terminal_type"`
  TerminalInfo      string                                          `json:"terminal_info"`
  ProdCode          string                                          `json:"prod_code"`
  NotifyUrl         string                                          `json:"notify_url"`
  ReturnUrl         string                                          `json:"return_url"`
}

func (this *AlipayOpenPublicMenuQueryRequest) GetApiMethodName() string {
  return "alipay.open.public.menu.query"
}

func (this *AlipayOpenPublicMenuQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicMenuQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicMenuQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicMenuQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicMenuQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicMenuQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicMenuQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicMenuQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
