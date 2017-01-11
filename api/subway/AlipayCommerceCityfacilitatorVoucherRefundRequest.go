package subway

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 地铁购票发码退款
// 地铁购票，商户发码后，调该接口从中间户退款
type AlipayCommerceCityfacilitatorVoucherRefundRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                      `json:"terminal_type"`
  TerminalInfo      string                                                      `json:"terminal_info"`
  ProdCode          string                                                      `json:"prod_code"`
  NotifyUrl         string                                                      `json:"notify_url"`
  ReturnUrl         string                                                      `json:"return_url"`
  BizContent        AlipayCommerceCityfacilitatorVoucherRefundRequestBizContent `json:"biz_content"`
}

type AlipayCommerceCityfacilitatorVoucherRefundRequestBizContent struct {
  CityCode  string `json:"city_code"` // 城市编码请参考查询：http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201504/t20150415_712722.html； 已支持城市：广州 440100，深圳 440300，杭州330100。
  TradeNo   string `json:"trade_no"`  // 支付宝交易号
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetApiMethodName() string {
  return "alipay.commerce.cityfacilitator.voucher.refund"
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayCommerceCityfacilitatorVoucherRefundRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
