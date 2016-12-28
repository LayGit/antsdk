package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradeCancelResponse struct {
  Result  AlipayTradeCancelResult `json:"alipay_trade_cancel_response"`
  Sign    string                  `json:"sign"`
}

type AlipayTradeCancelResult struct {
  api.CommonResponse
  TradeNo     string `json:"trade_no"`      // 支付宝交易号
  OutTradeNo  string `json:"out_trade_no"`  // 商户订单号
  RetryFlag   string `json:"retry_flag"`    // 是否需要重试
  Action      string `json:"action"`        // 本次撤销触发的交易动作 close：关闭交易，无退款 refund：产生了退款
}
