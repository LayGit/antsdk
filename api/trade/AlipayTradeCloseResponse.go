package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradeCloseResponse struct {
  api.AlipayResponse
  TradeNo     string `json:"trade_no"`      // 支付宝交易号
  OutTradeNo  string `json:"out_trade_no"`  // 创建交易传入的商户订单号
}
