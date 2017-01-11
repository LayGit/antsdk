package coupon

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingToolPointsQueryResponse struct {
  api.AlipayResponse
  AvailablePoints int `json:"available_points,string"`  // 可用集点
  FreezedPoints   int `json:"freezed_points,string"`    // 冻结集点
  TotalPoints     int `json:"total_points,string"`      // 累计集点
}
