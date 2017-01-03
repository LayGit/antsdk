package zhima

import (
  "github.com/LayGit/antsdk/api"
)

type ZhimaCustomerCertificationCertifyResponse struct {
  api.AlipayResponse
  BizNo         string `json:"biz_no"`          // 一次认证的唯一标识,在商户调用认证初始化接口的时候获取,认证完成返回的biz_no和请求的一致
  Passed        string `json:"passed"`          // 认证是否通过,通过为true,不通过为false
  FailedReason  string `json:"failed_reason"`   // 如果认证没有通过会显示失败原因,如果需要详细的失败原因请使用认证查询接口
}
