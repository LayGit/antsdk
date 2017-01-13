package data

import (
  "github.com/LayGit/antsdk/api"
)

type KoubeiMarketingDataCustomreportQueryResponse struct {
  api.AlipayResponse
  Count       int               `json:"count,string"` // 数据量
  ReportData  []ReportDataItem  `json:"report_data"`  // 满足自定义报表规则的报表数据
}

type ReportDataItem struct {
  RowData string `json:"row_data"`  // 表示一行数据，每个对象是一列的数据
}
