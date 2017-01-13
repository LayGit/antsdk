package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 会员卡查询
// 根据卡号或者持卡人信息查询会员卡信息
type AlipayMarketingCardQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                     `json:"terminal_type"`
  TerminalInfo      string                                     `json:"terminal_info"`
  ProdCode          string                                     `json:"prod_code"`
  NotifyUrl         string                                     `json:"notify_url"`
  ReturnUrl         string                                     `json:"return_url"`
  BizContent        AlipayMarketingCardQueryRequestBizContent  `json:"biz_content"`
}

type AlipayMarketingCardQueryRequestBizContent struct {
  TargetCardNo      string        `json:"target_card_no"`       // 操作卡号。 target_card_no为业务卡号，由开卡流程中，支付宝返回的业务卡号
  TargetCardNoType  string        `json:"target_card_no_type"`  // 卡号ID类型（会员卡查询，只能提供支付宝端的卡号） BIZ_CARD：支付宝卡号 D_QR_CODE：动态二维码（业务卡号对应的） D_BAR_CODE：动态条码（业务卡号对应的） 如果卡号不空，则类型不能为空
  CardUserInfo      CardUserInfo  `json:"card_user_info"`       // 用户信息 填写则作为附加条件查询
  ExtInfo           string        `json:"ext_info"`             // 扩展信息，暂时没有
}

func (this *AlipayMarketingCardQueryRequest) GetApiMethodName() string {
  return "alipay.marketing.card.query"
}

func (this *AlipayMarketingCardQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCardQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCardQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCardQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCardQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCardQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCardQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCardQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
