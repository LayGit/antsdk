package trade

type AlipayTradeCloseResponse struct {
  Result  AlipayTradeCloseResult  `json:"alipay_trade_close_response"`
  Sign    string                  `json:"sign"`
}

type AlipayTradeCloseResult struct {
  TradeNo     string `json:"trade_no"`      // 支付宝交易号
  OutTradeNo  string `json:"out_trade_no"`  // 创建交易传入的商户订单号
}
