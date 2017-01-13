package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicAccountResetResponse struct {
  api.AlipayResponse
  AgreementId string `json:"agreement_id"`  // 重置后的协议号，商户会员在支付宝服务窗账号中的唯一标识
}
