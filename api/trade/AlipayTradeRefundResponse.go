package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradeRefundResponse struct {
  api.AlipayResponse
  TradeNo               string            `json:"trade_no"`                 // 支付宝交易号
  OutTradeNo            string            `json:"out_trade_no"`             // 商户订单号
  OpenId                string            `json:"open_id"`                  // 买家支付宝用户号，该参数已废弃，请不要使用
  BuyerLogonId          string            `json:"buyer_logon_id"`           // 用户的登录id
  FundChange            string            `json:"fund_change"`              // 本次退款是否发生了资金变化
  RefundFee             string            `json:"refund_fee"`               // 退款总金额
  GMTRefundPay          string            `json:"gmt_refund_pay"`           // 退款支付时间
  RefundDetailItemList  []TradeFundBill   `json:"refund_detail_item_list"`  // 退款使用的资金渠道
  StoreName             string            `json:"store_name"`               // 交易在支付时候的门店名称
  BuyerUserId           string            `json:"buyer_user_id"`            // 买家在支付宝的用户id
}
