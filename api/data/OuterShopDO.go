package data

type OuterShopDO struct {
  ShopId  string `json:"shop_id"`   // 支付宝店铺ID
  OuterId string `json:"outer_id"`  // 合作商户ID
  Type    string `json:"type"`      // 合作商户类型 （mike、_2dFire） iSV自己定义自己的类型，推荐使用自己的域名
}
