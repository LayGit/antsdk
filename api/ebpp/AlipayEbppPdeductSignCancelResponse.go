package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppPdeductSignCancelResponse struct {
  api.AlipayResponse
  AgreementId     string `json:"agreement_id"`      // 支付宝代扣协议ID
  OutAgreementId  string `json:"out_agreement_id"`  // 商户代扣协议ID
  AgreementStatus string `json:"agreement_status"`  // 支付宝协议状态。解约成功则返回success
}
