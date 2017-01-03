package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppPdeductSignQueryResponse struct {
  api.AlipayResponse
  AgreementId     string `json:"agreement_id"`      // 协议ID
  SignDate        string `json:"sign_date"`         // 签约时间
  ChargeInst      string `json:"charge_inst"`       // 出账机构
  UserId          string `json:"user_id"`           // 用户ID
  BillKey         string `json:"bill_key"`          // 户号
  OutAgreementId  string `json:"out_agreement_id"`  // 朗新协议ID
}
