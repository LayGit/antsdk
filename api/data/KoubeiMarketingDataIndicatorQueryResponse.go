package data

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingDataIndicatorQueryResponse struct {
  api.AlipayResponse
  TotalSize       int     `json:"total_size,string"`  // 总条目数,供计算分页信息使用
  IndicatorInfos  string  `json:"indicator_infos"`    // SON格式数组，每个对象表示一个门店某个具体日期的指标信息，KEY为指标代码，VALUE为该指标对应的值,各biz_type入参以及返回值的详细信息参见快速接入
}
