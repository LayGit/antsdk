package coupon

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayPassInstanceAddResponse struct {
  api.AlipayResponse
  Success string `json:"success"`   // 操作成功标识【true：成功；false：失败】
  Result  string `json:"result"`    // 接口调用返回结果信息
}
