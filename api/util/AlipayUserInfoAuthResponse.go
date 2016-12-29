package util

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayUserInfoAuthResponse struct {
  Result  AlipayUserInfoAuthResult  `json:"alipay_user_info_auth_response"`
  Sign    string                    `json:"sign"`
}

type AlipayUserInfoAuthResult struct {
  api.CommonResponse
}
