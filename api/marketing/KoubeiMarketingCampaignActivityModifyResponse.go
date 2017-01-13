package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingCampaignActivityModifyResponse struct {
  api.AlipayResponse
  CampStatus  string `json:"camp_status"`   // 活动状态
  AuditStatus string `json:"audit_status"`  // 活动子状态，如审核中
}
