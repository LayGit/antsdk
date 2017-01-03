package util

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 查询某个应用授权AppAuthToken的授权信息
// 当商户把服务窗、店铺等接口的权限授权给ISV之后，支付宝会给ISV颁发一个app_auth_token。如若授权成功之后，ISV想知道用户的授权信息，如授权者、授权接口列表等信息，可以调用本接口查询某个app_auth_token对应的授权信息
type AlipayOpenAuthTokenAppQueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                                   `json:"terminal_type"`
  TerminalInfo  string                                   `json:"terminal_info"`
  ProdCode      string                                   `json:"prod_code"`
  NotifyUrl     string                                   `json:"notify_url"`
  ReturnUrl     string                                   `json:"return_url"`
  BizContent  AlipayOpenAuthTokenAppQueryRequestBizContent  `json:"biz_content"`
}

type AlipayOpenAuthTokenAppQueryRequestBizContent struct {
  AppAuthToken  string  `json:"app_auth_token"` // 应用授权令牌
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetApiMethodName() string {
  return "alipay.open.auth.token.app.query"
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenAuthTokenAppQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
