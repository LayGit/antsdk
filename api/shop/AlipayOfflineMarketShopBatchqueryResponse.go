package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMarketShopBatchqueryResponse struct {
  api.AlipayResponse
  ShopIds         []string  `json:"shop_ids"`         // 门店列表ID，逗号分隔
  CurrentPageNo   string    `json:"current_pageno"`   // 当前页码
  TotalPageno     string    `json:"total_pageno"`     // 总页码数目
}
