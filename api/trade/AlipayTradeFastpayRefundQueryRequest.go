package trade

// 统一收单交易退款查询
// 商户可使用该接口查询自已通过alipay.trade.refund提交的退款请求是否执行成功。
type AlipayTradeFastpayRefundQueryRequest struct {
  TradeNo       string `json:"trade_no"`        // 支付宝交易号，和商户订单号不能同时为空
  OutTradeNo    string `json:"out_trade_no"`    // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
  OutRequestNo  string `json:"out_request_no"`  // 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetMethod() string {
  return "alipay.trade.fastpay.refund.query"
}
