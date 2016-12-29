package trade

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 统一收单交易撤销接口
// 支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。 注意：只有发生支付系统超时或者支付结果未知时可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。提交支付交易后调用【查询订单API】，没有明确的支付结果再调用【撤销订单API】。
type AlipayTradeCancelRequest struct {
  api.IAlipayRequest
  TerminalType  string                              `json:"terminal_type"`
  TerminalInfo  string                              `json:"terminal_info"`
  ProdCode      string                              `json:"prod_code"`
  NotifyUrl     string                              `json:"notify_url"`
  ReturnUrl     string                              `json:"return_url"`
  BizContent    AlipayTradeCancelRequestBizContent  `json:"biz_content"`
}

type AlipayTradeCancelRequestBizContent struct {
  OutTradeNo  string `json:"out_trade_no"`  // 原支付请求的商户订单号,和支付宝交易号不能同时为空
  TradeNo     string `json:"trade_no"`      // 支付宝交易号，和商户订单号不能同时为空
}

func (this *AlipayTradeCancelRequest) GetApiMethodName() string {
  return "alipay.trade.cancel"
}

func (this *AlipayTradeCancelRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayTradeCancelRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayTradeCancelRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayTradeCancelRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayTradeCancelRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayTradeCancelRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayTradeCancelRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayTradeCancelRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
