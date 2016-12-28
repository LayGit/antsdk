package trade

// 优惠券信息
type VoucherDetail struct {
  Id                  string  `json:"id"`                          // 券id
  Name                string  `json:"name"`                        // 券名称
  Type                string  `json:"type"`                        // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券 ALIPAY_DISCOUNT_VOUCHER - 折扣券 ALIPAY_ITEM_VOUCHER - 单品优惠 注：不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码
  Amount              float64 `json:"amount,string"`               // 优惠券面额，它应该会等于商家出资加上其他出资方出资
  MerchantContribute  float64 `json:"merchant_contribute,string"`  // 商家出资（特指发起交易的商家出资金额）
  OtherContribute     float64 `json:"other_contribute,string"`     // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
  Memo                string  `json:"memo"`                        // 优惠券备注信息
}
