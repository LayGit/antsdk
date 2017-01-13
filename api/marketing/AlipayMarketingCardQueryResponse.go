package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayMarketingCardQueryResponse struct {
  api.AlipayResponse
  CardInfo  MerchantCard  `json:"card_info"`    // 商户卡信息
  SchemaUrl string        `json:"schema_url"`   // 商户会员卡页面跳转到支付宝卡券详情页面的schema地址
}
