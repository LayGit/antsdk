package util

import (
  "github.com/LayGit/antsdk/api"
)

type AlipaySystemOauthTokenResponse struct {
  api.AlipayResponse
  AlipayUserId  string `json:"alipay_user_id"`
  UserId        string `json:"user_id"`
  AccessToken   string `json:"access_token"`
  ExpiresIn     string `json:"expires_in"`
  RefreshToken  string `json:"refresh_token"`
  ReExpiresIn   string `json:"re_expires_in"`
}
