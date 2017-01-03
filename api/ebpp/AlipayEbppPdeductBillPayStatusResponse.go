package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppPdeductBillPayStatusResponse struct {
  api.AlipayResponse
  Status      string `json:"status"`          // 支付宝订单支付状态。 0：未知状态。 1：支付成功。 2：支付失败。
  AgreementId string  `json:"agreement_id"`   // 支付宝协议流水
  OutOrderNo  string  `json:"out_order_no"`   // 外部订单流水
  OrderNo     string  `json:"order_no"`       // 支付宝流billNo
}
