package data

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingCampaignCrowdBatchqueryResponse struct {
  api.AlipayResponse
  TotalNumber     int     `json:"total_number,string"`  // 返回接记录的总条数
  CrowdGroupSets  string  `json:"crowd_group_sets"`     // 人群组的基本信息，id表示人群分组的ID，name表示人群分组的名称，status表示人群分组的状态，目前只有status=ENABLE有效状态才返回，已经删除的为DISABLE的不返回
}
