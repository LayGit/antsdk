package data

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingCampaignTagsQueryResponse struct {
  api.AlipayResponse
  Tags string `json:"tags"` // 查询成功时返回人群标签信息查询失败时为空 code:表示标签code name:表示标签名称 valueRange：表示标签的取值范围 value:表示标签具体取值 label:描述信息 标签相关的详细信息参见附件。
}
