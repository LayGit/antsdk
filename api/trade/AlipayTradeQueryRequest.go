package trade

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 统一收单线下交易查询
// 该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。 需要调用查询接口的情况： 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知； 调用支付接口后，返回系统错误或未知交易状态情况； 调用alipay.trade.pay，返回INPROCESS的状态； 调用alipay.trade.cancel之前，需确认支付状态；
type AlipayTradeQueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                              `json:"terminal_type"`
  TerminalInfo  string                              `json:"terminal_info"`
  ProdCode      string                              `json:"prod_code"`
  NotifyUrl     string                              `json:"notify_url"`
  ReturnUrl     string                              `json:"return_url"`
  BizContent    AlipayTradeQueryRequestBizContent   `json:"biz_content"`
}

type AlipayTradeQueryRequestBizContent struct {
  OutTradeNo  string  `json:"out_trade_no"` // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
  TradeNo     string  `json:"trade_no"`     // 支付宝交易号，和商户订单号不能同时为空
}

func (this *AlipayTradeQueryRequest) GetApiMethodName() string {
  return "alipay.trade.query"
}

func (this *AlipayTradeQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayTradeQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayTradeQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayTradeQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayTradeQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayTradeQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayTradeQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayTradeQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
