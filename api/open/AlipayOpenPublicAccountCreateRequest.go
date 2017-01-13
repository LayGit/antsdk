package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 添加绑定商户会员号
// 当用户成为商户的关注用户后，可以在商户的服务窗中点击“添加绑定商户会员号”功能，支付宝系统收到操作请求后将该动作通知给商户（使用用户发送信息到商户接口，其中eventType（事件类型）为click，actionParam（按钮标识）为authentication），商户根据此通知调用商户回复消息接口（其中须包含Url链接地址），支付宝收到商户的回复消息中的链接地址后，自动跳转至商户平台的上商户会员绑定界面中，让用户完成账户绑定。 当用户有效完成账户绑定后，商户调用本接口，把绑定结果数据通知给支付宝。
type AlipayOpenPublicAccountCreateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                          `json:"terminal_type"`
  TerminalInfo      string                                          `json:"terminal_info"`
  ProdCode          string                                          `json:"prod_code"`
  NotifyUrl         string                                          `json:"notify_url"`
  ReturnUrl         string                                          `json:"return_url"`
  BizContent        AlipayOpenPublicAccountCreateRequestBizContent  `json:"biz_content"`
}

type AlipayOpenPublicAccountCreateRequestBizContent struct {
  BindAccountNo string `json:"bind_account_no"`   // 绑定帐号，建议在开发者的系统中保持唯一性
  DisplayName   string `json:"display_name"`      // 开发者期望在服务窗首页看到的关于该用户的显示信息，最长10个字符
  AgreementId   string `json:"agreement_id"`      // 账户添加成功，在支付宝与其对应的协议号。如果账户重复添加，接口保证幂等依然视为添加成功，返回此前该账户在支付宝对应的协议号。其他异常该字段不存在。
  RealName      string `json:"real_name"`         // 要绑定的商户会员的真实姓名，最长10个汉字
  FromUserId    string `json:"from_user_id"`      // 要绑定的商户会员对应的支付宝userid，2088开头长度为16位的字符串
  Remark        string `json:"remark"`            // 备注信息，开发者可以通过该字段纪录其他的额外信息
}

func (this *AlipayOpenPublicAccountCreateRequest) GetApiMethodName() string {
  return "alipay.open.public.account.create"
}

func (this *AlipayOpenPublicAccountCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicAccountCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicAccountCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicAccountCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicAccountCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicAccountCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicAccountCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicAccountCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
