package coupon

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingToolPrizesendAuthResponse struct {
  api.AlipayResponse
  Token string `json:"token"` // 发奖token，用于校验商户是否有权限给制定用户发奖
}
