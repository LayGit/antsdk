package ebpp

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 直连代扣协议查询接口
// 提供给朗新查询代扣协议状态的接口，用于进行双边对账，解决单边协议问题
type AlipayEbppPdeductSignQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                         `json:"terminal_type"`
  TerminalInfo      string                                         `json:"terminal_info"`
  ProdCode          string                                         `json:"prod_code"`
  NotifyUrl         string                                         `json:"notify_url"`
  ReturnUrl         string                                         `json:"return_url"`
  BizContent        AlipayEbppPdeductSignQueryRequestBizContent    `json:"biz_content"`
}

type AlipayEbppPdeductSignQueryRequestBizContent struct {
  UserId      string `json:"user_id"`       // 用户ID
  AgreementId string `json:"agreement_id"`  // 支付宝代扣协议Id。若协议id不传递，则需要保证业务类型、子业务类型、出账机构、户号必传
  BizType     string `json:"biz_type"`      // 业务类型。 JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用 WUYE：缴物业费 HK：信用卡还款 TX：手机充值
  SubBizType  string `json:"sub_biz_type"`  // 业务子类型。 WATER：缴水费 ELECTRIC：缴电费 GAS：缴燃气费 COMMUN：缴固话宽带 CATV：缴有线电视费 TRAFFIC：缴交通罚款 WUYE：缴物业费 HK：信用卡还款 CZ：手机充值
  ChargeInst  string `json:"charge_inst"`   // 支付宝缴费系统中的出账机构ID
  BillKey     string `json:"bill_key"`      // 户号，机构针对于每户的水、电都会有唯一的标识户号
}

func (this *AlipayEbppPdeductSignQueryRequest) GetApiMethodName() string {
  return "alipay.ebpp.pdeduct.sign.query"
}

func (this *AlipayEbppPdeductSignQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEbppPdeductSignQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEbppPdeductSignQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEbppPdeductSignQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEbppPdeductSignQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEbppPdeductSignQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEbppPdeductSignQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEbppPdeductSignQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
