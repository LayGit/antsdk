package zhima

import (
  "github.com/LayGit/antsdk/api"
)

type ZhimaCreditScoreBriefGetResponse struct {
  api.AlipayResponse
  BizNo         string `json:"biz_no"`          // 芝麻信用对于每一次请求返回的业务号。后续可以通过此业务号进行对账
  IsAdmittance  string `json:"is_admittance"`   // 准入判断结果 Y=准入,N=不准入,N/A=无法评估该用户的信用
}
