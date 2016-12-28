package trade

type AlipayTradePayResponse struct {
  Result  AlipayTradePayResult  `json:"alipay_trade_pay_response"`
  Sign    string                `json:"sign"`
}

type AlipayTradePayResult struct {
  TradeNo               string            `json:"trade_no"`                 // 支付宝交易号
  OutTradeNo            string            `json:"out_trade_no"`             // 商户订单号
  BuyerLogonId          string            `json:"buyer_logon_id"`           // 买家支付宝账号
  TotalAmount           float64           `json:"total_amount,string"`      // 交易金额
  ReceiptAmount         float64           `json:"receipt_amount,string"`    // 实收金额
  BuyerPayAmount        float64           `json:"buyer_pay_amount,string"`  // 买家付款的金额
  PointAmount           float64           `json:"point_amount,string"`      // 使用积分宝付款的金额
  InvoiceAmount         float64           `json:"invoice_amount,string"`    // 交易中可给用户开具发票的金额
  GMTPayment            string            `json:"gmt_payment"`              // 交易支付时间
  FundBillList          []TradeFundBill   `json:"fund_bill_list"`           // 交易支付使用的资金渠道
  CardBalance           float64           `json:"card_balance,string"`      // 支付宝卡余额
  StoreName             string            `json:"store_name"`               // 发生支付交易的商户门店名称
  BuyerUserId           string            `json:"buyer_user_id"`            // 买家在支付宝的用户id
  DiscountGoodsDetail   string            `json:"discount_goods_detail"`    // 本次交易支付所使用的单品券优惠的商品优惠信息
  VoucherDetailList     []VoucherDetail   `json:"voucher_detail_list"`      // 本交易支付时使用的所有优惠券信息
}
