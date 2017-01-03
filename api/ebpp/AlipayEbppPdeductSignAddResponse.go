package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppPdeductSignAddResponse struct {
  api.AlipayResponse
  AgreementId     string `json:"agreement_id"`      // 支付宝代扣协议ID
  SignDate        string `json:"sign_date"`         // 签约时间
  PayConfig       string `json:"pay_config"`        // 支付方式设置
  NotifyConfig    string `json:"notify_config"`     // 通知方式设置。
  OutAgreementId  string `json:"out_agreement_id"`  // 商户生成的代扣协议ID
  AgreementStatus string `json:"agreement_status"`  // 支付宝协议状态。签约成功则返回success
  ExtendField     string `json:"extend_field"`      // 扩展参数，可为空
}
