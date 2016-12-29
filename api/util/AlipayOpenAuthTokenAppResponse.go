package util

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenAuthTokenAppResponse struct {
  Result  AlipayOpenAuthTokenAppResult `json:"alipay_open_auth_token_app_response"`
  Sign    string                       `json:"sign"`
}

type AlipayOpenAuthTokenAppResult struct {
  api.CommonResponse
  UserId          string `json:"user_id"`           // 授权商户的user_id
  AuthAppId       string `json:"auth_app_id"`       // 授权商户的appid
  AppAuthToken    string `json:"app_auth_token"`    // 应用授权令牌
  AppRefreshToken string `json:"app_refresh_token"` // 刷新令牌
  ExpiresIn       string `json:"expires_in"`        // 应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒
  ReExpiresIn     string `json:"re_expires_in"`     // 刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒
}
