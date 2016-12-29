package trade

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 统一收单交易关闭接口
// 用于交易创建后，用户在一定时间内未进行支付，可调用该接口直接将未付款的交易进行关闭。
type AlipayTradeCloseRequest struct {
  api.IAlipayRequest
  TerminalType  string                              `json:"terminal_type"`
  TerminalInfo  string                              `json:"terminal_info"`
  ProdCode      string                              `json:"prod_code"`
  NotifyUrl     string                              `json:"notify_url"`
  ReturnUrl     string                              `json:"return_url"`
  BizContent    AlipayTradeCloseRequestBizContent   `json:"biz_content"`
}

type AlipayTradeCloseRequestBizContent struct {
  TradeNo     string `json:"trade_no"`      // 该交易在支付宝系统中的交易流水号。最短 16 位，最长 64 位。和out_trade_no不能同时为空，如果同时传了 out_trade_no和 trade_no，则以 trade_no为准。
  OutTradeNo  string `json:"out_trade_no"`  // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
  OperatorId  string `json:"operator_id"`   // 卖家端自定义的的操作员 ID
}

func (this AlipayTradeCloseRequest) GetApiMethodName() string {
  return "alipay.trade.close"
}

func (this *AlipayTradeCloseRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayTradeCloseRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayTradeCloseRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayTradeCloseRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayTradeCloseRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayTradeCloseRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayTradeCloseRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayTradeCloseRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
