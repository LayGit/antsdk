package subway

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 地铁购票核销码发码
// 商户APP调快捷支付付款后，获取地铁购票的核销码
type AlipayCommerceCityfacilitatorVoucherGenerateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                        `json:"terminal_type"`
  TerminalInfo      string                                                        `json:"terminal_info"`
  ProdCode          string                                                        `json:"prod_code"`
  NotifyUrl         string                                                        `json:"notify_url"`
  ReturnUrl         string                                                        `json:"return_url"`
  BizContent        AlipayCommerceCityfacilitatorVoucherGenerateRequestBizContent `json:"biz_content"`
}

type AlipayCommerceCityfacilitatorVoucherGenerateRequestBizContent struct {
  CityCode    string  `json:"city_code"`    // 市编码请参考查询：http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201504/t20150415_712722.html； 已支持城市：广州 440100，深圳 440300，杭州330100。
  TradeNo     string  `json:"trade_no"`     // 支付宝交易号（交易支付时，必须通过指定sellerId：2088121612215201，将钱支付到指定的中间户中）
  TotalFee    float64 `json:"total_fee"`    // 订单总金额，元为单位
  TicketNum   int     `json:"ticket_num"`   // 地铁票购票数量
  TicketType  string  `json:"ticket_type"`  // 地铁票种类
  SiteBegin   string  `json:"site_begin"`   // 起点站站点编码
  SiteEnd     string  `json:"site_end"`     // 终点站站点编码
  TicketPrice string  `json:"ticket_price"` // 单张票价，元为单价
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetApiMethodName() string {
  return "alipay.commerce.cityfacilitator.voucher.generate"
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayCommerceCityfacilitatorVoucherGenerateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
