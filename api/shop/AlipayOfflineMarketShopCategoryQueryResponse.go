package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMarketShopCategoryQueryResponse struct {
  api.AlipayResponse
  ShopCategoryConfigInfos []ShopCategoryConfigInfo `json:"shop_category_config_infos"`  // 门店类目配置信息，包括能够开店的叶子节点类目信息，以及类目约束配置信息。
}

type ShopCategoryConfigInfo struct {
  Id      string `json:"id"`        // 类目ID
  Name    string `json:"nm"`        // 类目名称
  Level   string `json:"level"`     // 类目层级
  Link    string `json:"link"`      // 类目层级路径
  IsLeaf  string `json:"is_leaf"`   // 是否是叶子节点
}
