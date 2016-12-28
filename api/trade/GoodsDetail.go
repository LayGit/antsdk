package trade

type GoodsDetail struct {
  GoodsId       string  `json:"goods_id"`         // 商品的编号
  AlipayGoodsId string  `json:"alipay_goods_id"`  // 支付宝定义的统一商品编号
  GoodsName     string  `json:"goods_name"`       // 商品名称
  Quantity      int     `json:"quantity,string"`  // 商品数量
  Price         float64 `json:"price,string"`     // 商品单价，单位为元
  GoodsCategory string  `json:"goods_category"`   // 商品类目
  Body          string  `json:"body"`             // 商品描述信息
  ShowURL       string  `json:"show_url"`         // 商品的展示地址
}
