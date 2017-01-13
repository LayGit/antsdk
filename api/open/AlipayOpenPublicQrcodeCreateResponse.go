package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicQrcodeCreateResponse struct {
  api.AlipayResponse
  CodeImg       string `json:"code_img"`      // 二维码图片地址，可跳转到实际图片
  ExpireSecond  string `json:"expire_second"` // 二维码有效时间，单位（秒）。永久码暂时忽略该信息
}
