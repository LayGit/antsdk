package trade

// 统一收单交易结算接口
// 用于在线下场景交易支付后，进行结算
type AlipayTradeOrderSettleRequest struct {
  OutRequestNo      string                          `json:"out_request_no"`     // 结算请求流水号
  TradeNo           string                          `json:"trade_no"`           // 支付宝订单号
  RoyaltyParameters []OpenApiRoyaltyDetailInfoPojo  `json:"royalty_parameters"` // 分账明细信息
  OperatorId        string                          `json:"operator_id"`        // 操作员id
}

func (this *AlipayTradeOrderSettleRequest) GetMethod() string {
  return "alipay.trade.order.settle"
}
