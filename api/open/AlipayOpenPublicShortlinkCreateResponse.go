package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicShortlinkCreateResponse struct {
  api.AlipayResponse
  ShortLink string `json:"shortlink"` // 生成的带参推广短链接
}
