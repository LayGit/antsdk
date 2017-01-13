package marketing

type BenefitInfoDetail struct {
  BenefitType string `json:"benefit_type"`  // 权益类型 PRE_FUND（卡面额） DISCOUNT：折扣金额 COUPON：券
  Amount      string `json:"amount"`        // PRE_FUND：实际核销或者商户赠送的金额 DISCOUNT：实际折扣掉的金额（获取权益不支持该类型） COUPON：实际核销或者商户赠送的券
  Description string `json:"description"`   // 产生核销或者赠送权益的描述
  Count       string `json:"count"`         // COUPON：当核销或者赠送券时，可以设置该值
}
