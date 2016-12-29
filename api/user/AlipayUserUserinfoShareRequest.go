package user

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 支付宝钱包用户信息共享
// 外部应用上架到支付宝钱包，当支付宝用户从钱包访问外部应用时，会跳转到外部应用并带上用户的授权码。 外部应用用授权码调用授权令牌交换API（alipay.system.oauth.token）可得到授权令牌。 用授权令牌调用此接口得到支付宝会员相关信息。 特别说明：此接口的不需要授权是指不需外部应用主动引导用户授权，支付宝钱包会在引导用户授权后， 带上授权码再跳转到外部应用。
type AlipayUserUserinfoShareRequest struct {
  api.IAlipayRequest
  TerminalType  string                                    `json:"terminal_type"`
  TerminalInfo  string                                    `json:"terminal_info"`
  ProdCode      string                                    `json:"prod_code"`
  NotifyUrl     string                                    `json:"notify_url"`
  ReturnUrl     string                                    `json:"return_url"`
}

func (this *AlipayUserUserinfoShareRequest) GetApiMethodName() string {
  return "alipay.user.userinfo.share"
}

func (this *AlipayUserUserinfoShareRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayUserUserinfoShareRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayUserUserinfoShareRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayUserUserinfoShareRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayUserUserinfoShareRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayUserUserinfoShareRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayUserUserinfoShareRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayUserUserinfoShareRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
