package trade

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayTradeQueryResponse struct {
  api.AlipayResponse
  TradeNo             string            `json:"trade_no"`                 // 支付宝交易号
  OutTradeNo          string            `json:"out_trade_no"`             // 商家订单号
  OpenId              string            `json:"open_id"`                  // 买家支付宝用户号，该字段将废弃，不要使用
  BuyerLogonId        string            `json:"buyer_logon_id"`           // 买家支付宝账号
  TradeStatus         string            `json:"trade_status"`             // 交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
  TotalAmount         float64           `json:"total_amount,string"`      // 交易的订单金额，单位为元，两位小数。
  ReceiptAmount       float64           `json:"receipt_amount,string"`    // 实收金额，单位为元，两位小数。
  BuyerPayAmount      float64           `json:"buyer_pay_amount,string"`  // 买家实付金额，单位为元，两位小数。
  PointAmount         float64           `json:"point_amount,string"`      // 积分支付的金额，单位为元，两位小数。
  InvoiceAmount       float64           `json:"invoice_amount,string"`    // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
  SendPayDate         string            `json:"send_pay_date"`            // 本次交易打款给卖家的时间
  AlipayStoreId       string            `json:"alipay_store_id"`          // 支付宝店铺编号
  StoreId             string            `json:"store_id"`                 // 商户门店编号
  TerminalId          string            `json:"terminal_id"`              // 商户机具终端编号
  FundBillList        []TradeFundBill   `json:"fund_bill_list"`           // 交易支付使用的资金渠道
  StoreName           string            `json:"store_name"`               // 请求交易支付中的商户店铺的名称
  BuyerUserId         string            `json:"buyer_user_id"`            // 买家在支付宝的用户id
  DiscountGoodsDetail string            `json:"discount_goods_detail"`    // 本次交易支付所使用的单品券优惠的商品优惠信息
  IndustrySepcDetail  string            `json:"industry_sepc_detail"`     // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
}
