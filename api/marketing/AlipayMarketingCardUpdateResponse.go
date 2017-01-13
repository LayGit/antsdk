package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayMarketingCardUpdateResponse struct {
  api.AlipayResponse
  ResultCode string `json:"result_code"`  // 二级错误处理结果（如果公用返回结果为false，则可以看这个接口判断明细原因） 如果公用返回为true，则该字段为空
}
