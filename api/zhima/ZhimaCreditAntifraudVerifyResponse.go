package zhima

import (
  "github.com/LayGit/antsdk/api"
)

type ZhimaCreditAntifraudVerifyResponse struct {
  api.AlipayResponse
  BizNo       string `json:"biz_no"`        // 芝麻信用对于每一次请求返回的业务号。后续可以通过此业务号进行对账
  VerifyCode  string `json:"verify_code"`   // 风险因素code列表
}
