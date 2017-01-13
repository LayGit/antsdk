package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 口碑菜品热度查询
// 口碑菜品热度分页查询，并根据热度升降序
type AlipayOfflineProviderDishQueryRequest struct {
  api.IAlipayRequest
  TerminalType        string                                           `json:"terminal_type"`
  TerminalInfo        string                                           `json:"terminal_info"`
  ProdCode            string                                           `json:"prod_code"`
  NotifyUrl           string                                           `json:"notify_url"`
  ReturnUrl           string                                           `json:"return_url"`
  BizContent          AlipayOfflineProviderDishQueryRequestBizContent  `json:"biz_content"`
}

type AlipayOfflineProviderDishQueryRequestBizContent struct {
  ShopId        string  `json:"shop_id"`        // 外部商户ID（必填）需要查询的商家名下的门店/子商户的ID
  OuterDishId   string  `json:"outer_dish_id"`  // 菜品ID，当需要查询指定菜品的时候传递，非必填
  DishTypeName  string  `json:"dish_type_name"` // 商家分类名称，商家自定义的菜品分类的名称，非必填
  OrderBy       string  `json:"order_by"`       // order_by：1，菜品热度升序查询，order_by：2，菜品热度降序查询。非必填，为空时默认为2
  Page          int     `json:"page"`           // 需要查询的第几页信息。非必填。默认为1
  PageSize      int     `json:"page_size"`      // 分页查询每页的条数，默认为20条，非必填
}

func (this *AlipayOfflineProviderDishQueryRequest) GetApiMethodName() string {
  return "alipay.offline.provider.dish.query"
}

func (this *AlipayOfflineProviderDishQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineProviderDishQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineProviderDishQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineProviderDishQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineProviderDishQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineProviderDishQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineProviderDishQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineProviderDishQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
