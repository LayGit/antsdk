package trade

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 统一收单交易结算接口
// 用于在线下场景交易支付后，进行结算
type AlipayTradeOrderSettleRequest struct {
  api.IAlipayRequest
  TerminalType  string                                   `json:"terminal_type"`
  TerminalInfo  string                                   `json:"terminal_info"`
  ProdCode      string                                   `json:"prod_code"`
  NotifyUrl     string                                   `json:"notify_url"`
  ReturnUrl     string                                   `json:"return_url"`
  BizContent    AlipayTradeOrderSettleRequestBizContent  `json:"biz_content"`
}

type AlipayTradeOrderSettleRequestBizContent struct {
  OutRequestNo      string                          `json:"out_request_no"`     // 结算请求流水号
  TradeNo           string                          `json:"trade_no"`           // 支付宝订单号
  RoyaltyParameters []OpenApiRoyaltyDetailInfoPojo  `json:"royalty_parameters"` // 分账明细信息
  OperatorId        string                          `json:"operator_id"`        // 操作员id
}

func (this *AlipayTradeOrderSettleRequest) GetApiMethodName() string {
  return "alipay.trade.order.settle"
}

func (this *AlipayTradeOrderSettleRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayTradeOrderSettleRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayTradeOrderSettleRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayTradeOrderSettleRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayTradeOrderSettleRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayTradeOrderSettleRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayTradeOrderSettleRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayTradeOrderSettleRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
