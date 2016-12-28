package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradeOrderSettleResponse struct {
  Result  AlipayTradeOrderSettleResult  `json:"alipay_trade_order_settle_response"`
  Sign    string                        `json:"sign"`
}

type AlipayTradeOrderSettleResult struct {
  api.CommonResponse
  TradeNo string `json:"trade_no"` //支付宝交易号
}
