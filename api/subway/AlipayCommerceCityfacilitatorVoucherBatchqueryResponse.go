package subway

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayCommerceCityfacilitatorVoucherBatchqueryResponse struct {
  api.AlipayResponse
  Tickets []TicketDetailInfo `json:"tickets"` // 查询到的订单信息列表
}

type TicketDetailInfo struct {
  TradeNo           string  `json:"trade_no"`           // 支付宝交易号
  Amount            float64 `json:"amount"`             // 总金额，元为单位
  StartStation      string  `json:"start_station"`      // 起点站编码
  EndStation        string  `json:"end_station"`        // 终点站编码
  Quantity          string  `json:"quantity"`           // 票数量
  Status            string  `json:"status"`             // 订单状态
  TicketPrice       float64 `json:"ticket_price"`       // 单价，元为单位
  StartStationName  string  `json:"start_station_name"` // 起点站中文名称
  EndStationName    string  `json:"end_station_name"`   // 终点站中文名称
  TicketType        string  `json:"ticket_type"`        // 票类型
}
