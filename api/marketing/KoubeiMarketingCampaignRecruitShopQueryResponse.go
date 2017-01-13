package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingCampaignRecruitShopQueryResponse struct {
  api.AlipayResponse
  ShopCount string        `json:"shop_count"`
  PlanId    string        `json:"plan_id"`
  ShopInfo  []RecruitInfo `json:"shop_info"`
}
