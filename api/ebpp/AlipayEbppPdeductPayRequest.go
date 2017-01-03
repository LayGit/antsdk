package ebpp

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 公共事业缴费直连代扣扣款支付接口
// 直连代扣机构根据用户个人签约协议，按期按账单请求从用户账户扣款的接口。
type AlipayEbppPdeductPayRequest struct {
  api.IAlipayRequest
  TerminalType      string  `json:"terminal_type"`
  TerminalInfo      string  `json:"terminal_info"`
  ProdCode          string  `json:"prod_code"`
  NotifyUrl         string  `json:"notify_url"`
  ReturnUrl         string  `json:"return_url"`
  AgreementId       string  `json:"agreement_id"`
  UserId            string  `json:"user_id"`
  BillKey           string  `json:"bill_key"`
  Pid               string  `json:"pid"`
  OutOrderNo        string  `json:"out_order_no"`
  PayAmount         float64 `json:"pay_amount"`
  FineAmount        float64 `json:"fine_amount"`
  BillDate          string  `json:"bill_date"`
  AgentChannel      string  `json:"agent_channel"`
  AgentCode         string  `json:"agent_code"`
  Memo              string  `json:"memo"`
  ExtendField       string  `json:"extend_field"`
}

func (this *AlipayEbppPdeductPayRequest) GetApiMethodName() string {
  return "alipay.ebpp.pdeduct.pay"
}

func (this *AlipayEbppPdeductPayRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEbppPdeductPayRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEbppPdeductPayRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEbppPdeductPayRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEbppPdeductPayRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEbppPdeductPayRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEbppPdeductPayRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEbppPdeductPayRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("agreement_id", this.AgreementId)
  txtParams.Put("user_id", this.UserId)
  txtParams.Put("bill_key", this.BillKey)
  txtParams.Put("pid", this.Pid)
  txtParams.Put("out_order_no", this.OutOrderNo)
  txtParams.Put("pay_amount", this.PayAmount)
  txtParams.Put("fine_amount", this.FineAmount)
  txtParams.Put("bill_date", this.BillDate)
  txtParams.Put("agent_channel", this.AgentChannel)
  txtParams.Put("agent_code", this.AgentCode)
  txtParams.Put("memo", this.Memo)
  txtParams.Put("extend_field", this.ExtendField)
  return txtParams
}
