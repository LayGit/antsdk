package data

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingDataCustomreportSaveResponse struct {
  api.AlipayResponse
  ConditionKey string `json:"condition_key"`  // 自定义报表的规则ID
}
