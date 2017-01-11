package finance

import (
  "github.com/LayGit/antsdk/api"
)

type MybankFinanceYulibaoPriceQueryResponse struct {
  api.AlipayResponse
  YlbPriceDetailInfos []YLBPriceDetailInfo `json:"ylb_price_detail_infos"`  // 余利宝行情信息列表
}

type YLBPriceDetailInfo struct {
  PriceDate           string  `json:"price_date"`            // 余利宝行情日期，格式为yyyy-MM-dd HH:mm:ss
  SevendaysYeildRate  float64 `json:"sevendays_yeild_rate,string"`  // 七日年化收益率（%），精确到小数点后4位
  TenthousandIncome   float64 `json:"tenthousand_income,string"`    // 万份收益金额，单位为元
}
