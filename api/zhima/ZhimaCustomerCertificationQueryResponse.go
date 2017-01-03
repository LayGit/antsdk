package zhima

import (
  "github.com/LayGit/antsdk/api"
)

type ZhimaCustomerCertificationQueryResponse struct {
  api.AlipayResponse
  Passed          string `json:"passed"`            // 认证是否通过,通过为true,不通过为false
  FailedReason    string `json:"failed_reason"`     // 如果认证没有通过会显示失败原因,更详细的情况在channel_statuses参数里面
  ChannelStatuses string `json:"channel_statuses"`  // 认证过程成认证过的各渠道的状态,中间材料等数据
}
