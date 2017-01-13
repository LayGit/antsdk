package data

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineProviderDishQueryResponse struct {
  api.AlipayResponse
  List      []IsvShopDishModel  `json:"list"`         // 菜品列表信息
  Items     int                 `json:"items"`        // 总共有多少条菜品信息。可用于计算分页。
  Page      int                 `json:"page"`         // 当前数据所在的页码数
  Pages     int                 `json:"pages"`        // 当前条件下查询结果总的页码数
  PageSize  int                 `json:"page_size"`    // 当前查询结果分页的条数，可用于计算分页
}
