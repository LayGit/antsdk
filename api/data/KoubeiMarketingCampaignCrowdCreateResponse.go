package data

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingCampaignCrowdCreateResponse struct {
  api.AlipayResponse
  CrowdGroupId string `json:"crowd_group_id"` // 返回的人群组的唯一标识
}
