package trade

// 统一收单线下交易查询
// 该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。 需要调用查询接口的情况： 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知； 调用支付接口后，返回系统错误或未知交易状态情况； 调用alipay.trade.pay，返回INPROCESS的状态； 调用alipay.trade.cancel之前，需确认支付状态；
type AlipayTradeQueryRequest struct {
  OutTradeNo  string  `json:"out_trade_no"` // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
  TradeNo     string  `json:"trade_no"`     // 支付宝交易号，和商户订单号不能同时为空
}

func (this *AlipayTradeQueryRequest) GetMethod() string {
  return "alipay.trade.query"
}
