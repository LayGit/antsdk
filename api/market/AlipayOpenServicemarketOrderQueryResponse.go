package market

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenServicemarketOrderQueryResponse struct {
  api.AlipayResponse
  OrderItems  []OrderItem `json:"order_items"`    // 订单明细列表
  TotalSize   int         `json:"total_size"`     // 总记录数
  CommodityId string      `json:"commodity_id"`   // 订购服务商品ID
  CurrentPage int         `json:"current_page"`   // 当前查询页（本接口支持最多查询100条记录）
  Status      string      `json:"status"`         // MERCHANT_ORDED（待服务商接单）
}
