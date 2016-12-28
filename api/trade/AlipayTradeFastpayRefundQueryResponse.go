package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradeFastpayRefundQueryResponse struct {
  Result  AlipayTradeFastpayRefundQueryResult `json:"alipay_trade_fastpay_refund_query_response"`
  Sign    string                              `json:"sign"`
}

type AlipayTradeFastpayRefundQueryResult struct {
  api.CommonResponse
  TradeNo       string  `json:"trade_no"`              // 支付宝交易号
  OutTradeNo    string  `json:"out_trade_no"`          // 创建交易传入的商户订单号
  OutRequestNo  string  `json:"out_request_no"`        // 本笔退款对应的退款请求号
  RefundReason  string  `json:"refund_reason"`         // 发起退款时，传入的退款原因
  TotalAmount   float64 `json:"total_amount,string"`   // 该笔退款所对应的交易的订单金额
  RefundAmount  float64 `json:"refund_amount,string"`  // 本次退款请求，对应的退款金额
}
