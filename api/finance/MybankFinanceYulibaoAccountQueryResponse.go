package finance

import (
  "github.com/LayGit/antsdk/api"
)

type MybankFinanceYulibaoAccountQueryResponse struct {
  api.AlipayResponse
  TotalAmount         float64             `json:"total_amount,string"`      // 余利宝总余额，单位为元
  AvailableAmount     float64             `json:"available_amount,string"`  // 可用份额，单位为元
  FreezeAmount        float64             `json:"freeze_amount,string"`     // 业务冻结份额，单位为元
  SysFreezeAmount     float64             `json:"sys_freeze_amount,string"` // 系统冻结份额，单位为元（建议不展示给用户）
  YlbProfitDetailInfo YLBProfitDetailInfo `json:"ylb_profit_detail_info"`   // 余利宝收益详情
}

type YLBProfitDetailInfo struct {
  DayProfit   float64 `json:"day_profit,string"`    // 近1日收益，单位为元
  WeekProfit  float64 `json:"week_profit,string"`   // 近1周收益，单位为元
  MonthProfit float64 `json:"month_profit,string"`  // 近1月收益，单位为元
  TotalProfit float64 `json:"total_profit"`         // 历史累计收益，单位为元
}
