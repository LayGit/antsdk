package marketing

type PubChannelDTO struct {
  PubChannel  string `json:"pub_channel"` // 1、SHOP_DETAIL:店铺详情页 2、PAYMENT_RESULT: 支付成功页（支付成功页暂不支持）
  ExtInfo     string `json:"ext_info"`    // 扩展信息，无需配置
}
