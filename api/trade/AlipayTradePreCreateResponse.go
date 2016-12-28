package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradePreCreateResponse struct {
  Result  AlipayTradePreCreateResult  `json:"alipay_trade_precreate_response"`
  Sign    string                      `json:"sign"`
}

type AlipayTradePreCreateResult struct {
  api.CommonResponse
  OutTradeNo  string  `json:"out_trade_no"` // 商户的订单号
  QRCode      string  `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
}
