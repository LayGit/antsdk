package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayMarketingCardOpenResponse struct {
  api.AlipayResponse
  CardInfo MerchantCard `json:"card_info"` // 商户卡信息（包括支付宝分配的业务卡号）
}
