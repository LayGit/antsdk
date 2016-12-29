package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMarketShopSummaryBatchqueryResponse struct {
  api.AlipayResponse
  PageSize          int                         `json:"page_size,string"`           // 每页记录数
  CurrentPageNo     int                         `json:"current_page_no,string"`     // 当前页码
  TotalPageNo       int                         `json:"total_page_no,string"`       // 总页码数目
  TotalItems        int                         `json:"total_items,string"`         // 总记录数
  ShopSummaryInfos  []ShopSummaryQueryResponse  `json:"shop_summary_infos"`         // 支付宝门店摘要信息列表
}

type ShopSummaryQueryResponse struct {
  ShopId          string              `json:"shop_id"`            // 门店ID
  MainShopName    string              `json:"main_shop_name"`     // 主门店名
  BranchShopName  string              `json:"branch_shop_name"`   // 分店名
  Status          string              `json:"status"`             // 门店状态，OPEN：营业中、PAUSE：暂停营业、FREEZE：已冻结、CLOSE:门店已关闭
  CategoryInfos   []ShopCategoryInfo  `json:"category_infos"`     // 门店类目列表
  BrandName       string              `json:"brand_name"`         // 品牌名，不填写则默认为其它品牌
  Address         string              `json:"address"`            // 门店地址
  ProvinceCode    string              `json:"province_code"`      // 省份编码，国标码，详见国家统计局数据
  CityCode        string              `json:"city_code"`          // 城市编码，国标码，详见国家统计局数据
  DistrictCode    string              `json:"district_code"`      // 区县编码，国标码，详见国家统计局数据
  BusinessTime    string              `json:"business_time"`      // 经营时间
  ShopType        string              `json:"shop_type"`          // COMMON（普通门店）、MALL（商圈）
  GMTCreate       string              `json:"gmt_create"`         // 创建时间
}

type ShopCategoryInfo struct {
  CategoryId    string `json:"category_id"`     // 类目编号
  CategoryName  string `json:"category_name"`   // 类目名称
  CategoryLevel string `json:"category_level"`  // 类目层级,目前最多支持1、2、3三级
}
