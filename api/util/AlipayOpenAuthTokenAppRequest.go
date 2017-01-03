package util

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 换取应用授权令牌
// 换取应用授权令牌。在应用授权的场景下，商户把名下应用授权给ISV后，支付宝会给ISV颁发应用授权码app_auth_code，ISV可通过获取到的app_auth_code换取app_auth_token。app_auth_code作为换取app_auth_token的票据，每次用户授权带上的app_auth_code将不一样，app_auth_code只能使用一次，一天（从当前时间算起的24小时）未被使用自动过期。 刷新应用授权令牌，ISV可通过获取到的refresh_token刷新app_auth_token，刷新后老的refresh_token会在一段时间后失效（失效时间为接口返回的re_expires_in）。
type AlipayOpenAuthTokenAppRequest struct {
  api.IAlipayRequest
  TerminalType  string                                   `json:"terminal_type"`
  TerminalInfo  string                                   `json:"terminal_info"`
  ProdCode      string                                   `json:"prod_code"`
  NotifyUrl     string                                   `json:"notify_url"`
  ReturnUrl     string                                   `json:"return_url"`
  BizContent    AlipayOpenAuthTokenAppRequestBizContent  `json:"biz_content"`
}

type AlipayOpenAuthTokenAppRequestBizContent struct {
  GrantType     string `json:"grant_type"`    // uthorization_code表示换取app_auth_token。 refresh_token表示刷新app_auth_token。
  Code          string `json:"code"`          // 授权码，如果grant_type的值为authorization_code。该值必须填写
  RefreshToken  string `json:"refresh_token"` // 刷新令牌，如果grant_type值为refresh_token。该值不能为空。该值来源于此接口的返回值app_refresh_token（至少需要通过grant_type=authorization_code调用此接口一次才能获取）
}

func (this *AlipayOpenAuthTokenAppRequest) GetApiMethodName() string {
  return "alipay.open.auth.token.app"
}

func (this *AlipayOpenAuthTokenAppRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenAuthTokenAppRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenAuthTokenAppRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenAuthTokenAppRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenAuthTokenAppRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenAuthTokenAppRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenAuthTokenAppRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenAuthTokenAppRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
