package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 重新设置绑定商家会员号
// 如果商户想要重置已经添加的外部账户，可以通过该接口完成。重置后，原有的外部户将删除，新的外部户添加进去。
type AlipayOpenPublicAccountResetRequest struct {
  api.IAlipayRequest
  TerminalType      string                                          `json:"terminal_type"`
  TerminalInfo      string                                          `json:"terminal_info"`
  ProdCode          string                                          `json:"prod_code"`
  NotifyUrl         string                                          `json:"notify_url"`
  ReturnUrl         string                                          `json:"return_url"`
  BizContent        AlipayOpenPublicAccountResetRequestBizContent   `json:"biz_content"`
}

type AlipayOpenPublicAccountResetRequestBizContent struct {
  BindAccountNo string `json:"bind_account_no"` // 绑定帐号，建议在开发者的系统中保持唯一性
  DisplayName   string `json:"display_name"`    // 开发者期望在服务窗首页看到的关于该用户的显示信息，最长10个字符
  AgreementId   string `json:"agreement_id"`    // 需要重置的协议号，商户会员在支付宝服务窗账号中的唯一标识
  RealName      string `json:"real_name"`       // 要绑定的商户会员的真实姓名，最长10个汉字
  FromUserId    string `json:"from_user_id"`    // 要绑定的商户会员对应的支付宝userid，2088开头长度为16位的字符串
  Remark        string `json:"remark"`          // 备注信息，开发者可以通过该字段纪录其他的额外信息
}

func (this *AlipayOpenPublicAccountResetRequest) GetApiMethodName() string {
  return "alipay.open.public.account.reset"
}

func (this *AlipayOpenPublicAccountResetRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicAccountResetRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicAccountResetRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicAccountResetRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicAccountResetRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicAccountResetRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicAccountResetRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicAccountResetRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
