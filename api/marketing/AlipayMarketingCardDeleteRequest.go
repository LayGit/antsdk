package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 会员卡删卡
// 通过API接口删除商户会员卡
type AlipayMarketingCardDeleteRequest struct {
  api.IAlipayRequest
  TerminalType      string                                     `json:"terminal_type"`
  TerminalInfo      string                                     `json:"terminal_info"`
  ProdCode          string                                     `json:"prod_code"`
  NotifyUrl         string                                     `json:"notify_url"`
  ReturnUrl         string                                     `json:"return_url"`
  BizContent        AlipayMarketingCardDeleteRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCardDeleteRequestBizContent struct {
  OutSerialNo       string `json:"out_serial_no"`       // 商户端删卡业务流水号（商户确保流水号唯一性）
  TargetCardNo      string `json:"target_card_no"`      // 支付宝业务卡号，开卡接口中返回获取
  TargetCardNoType  string `json:"target_card_no_type"` // 卡号ID类型 BIZ_CARD：支付宝卡号
  ReasonCode        string `json:"reason_code"`         // 删卡原因 USER_UNBUND：用户解绑（可以重新绑定） CANCEL：销户（完成销户后，就不能再重新绑定） PRESENT：转赠（可以重新绑定）
  ExtInfo           string `json:"ext_info"`            // 删卡扩展参数，json格式。 用于商户的特定业务信息的传递，只有商户与支付宝约定了传递此参数且约定了参数含义，此参数才有效。 目前支持如下key： new_card_no：新卡号 donee_user_id：受赠人userId
}

func (this *AlipayMarketingCardDeleteRequest) GetApiMethodName() string {
  return "alipay.marketing.card.delete"
}

func (this *AlipayMarketingCardDeleteRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCardDeleteRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCardDeleteRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCardDeleteRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCardDeleteRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCardDeleteRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCardDeleteRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCardDeleteRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
