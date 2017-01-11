package subway

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayCommerceCityfacilitatorStationQueryResponse struct {
  api.AlipayResponse
  SupportStarts []StationDetailInfo `json:"support_starts"` // 支持设为起点的站点列表
}

type StationDetailInfo struct {
  Name    string `json:"name"`      // 站点中文名称
  Code    string `json:"code"`      // 站点编码
  ExtCode string `json:"ext_code"`  // 站点外部编码
}
