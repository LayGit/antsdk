package trade

type TradeFundBill struct {
  FundChannel string  `json:"fund_channel"`        // 交易使用的资金渠道，详见 支付渠道列表(https://doc.open.alipay.com/doc2/detail?treeId=26&articleId=103259&docType=1)
  Amount      float64 `json:"amount,string"`       // 该支付工具类型所使用的金额
  RealAmount  float64 `json:"real_amount,string"`  // 渠道实际付款金额
}
