package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppPdeductPayResponse struct {
  api.AlipayResponse
  AgreementId   string `json:"agreement_id"`    // 支付宝代扣协议ID
  OutOrderNo    string `json:"out_order_no"`    // 商户代扣业务流水
  BillNo        string `json:"bill_no"`         // 支付宝订单流水号
  ResultStatus  string `json:"result_status"`   // 订单支付状态。 0：未知 1：成功 2：失败
  ExtendField   string `json:"extend_field"`    // 扩展参数
}
