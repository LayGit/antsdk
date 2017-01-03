package zhima

import (
  "github.com/LayGit/antsdk/api"
)

type ZhimaCustomerCertificationInitializeResponse struct {
  api.AlipayResponse
  BizNo string `json:"biz_no"` // 本次认证的唯一标识,商户需要记录
}
