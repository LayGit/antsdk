package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingCampaignActivityCreateResponse struct {
  api.AlipayResponse
  AuditStatus string `json:"audit_status"`  // 活动审批状态， 仅限服务商代商户创建活动时返回 AUDITING，审核中,REJECT为驳回，不返回表示通过
  CampId      string `json:"camp_id"`       // 活动ID
  CampStatus  string `json:"camp_status"`   // 活动状态，目前返回以下状态， STARTING，活动启动中 STARTED，活动已启动
}
