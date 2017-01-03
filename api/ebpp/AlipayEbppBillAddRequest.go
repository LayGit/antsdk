package ebpp

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 创建账单
type AlipayEbppBillAddRequest struct {
  api.IAlipayRequest
  TerminalType        string  `json:"terminal_type"`
  TerminalInfo        string  `json:"terminal_info"`
  ProdCode            string  `json:"prod_code"`
  NotifyUrl           string  `json:"notify_url"`
  ReturnUrl           string  `json:"return_url"`
  MerchantOrderNo     string  `json:"merchant_order_no"`      // 输出机构的业务流水号，需要保证唯一性
  OrderType           string  `json:"order_type"`             // 支付宝订单类型。公共事业缴纳JF,信用卡还款HK
  SubOrderType        string  `json:"sub_order_type"`         // 子业务类型是业务类型的下一级概念，例如：WATER表示JF下面的水费，ELECTRIC表示JF下面的电费，GAS表示JF下面的燃气费。
  ChargeInst          string  `json:"charge_inst"`            // 支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
  BillKey             string  `json:"bill_key"`               // 账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
  OwnerName           string  `json:"owner_name"`             // 拥有该账单的用户姓名
  PayAmount           float64 `json:"pay_amount"`             // 缴费金额。用户支付的总金额。单位为：RMB Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。
  ServiceAmount       float64 `json:"service_amount"`         // 账单的服务费。
  BillDate            string  `json:"bill_date"`              // 账单的账期，例如201203表示2012年3月的账单。
  Mobile              string  `json:"mobile"`                 // 用户的手机号
  TrafficLocation     string  `json:"traffic_location"`       // 交通违章地点，sub_order_type=TRAFFIC时填写。
  TrafficRegulations  string  `json:"traffic_regulations"`    // 违章行为，sub_order_type=TRAFFIC时填写。
  BankBillNo          string  `json:"bank_bill_no"`           // 外部订单号
  ExtendField         string  `json:"extend_field"`           // 扩展属性
}

func (this *AlipayEbppBillAddRequest) GetApiMethodName() string {
  return "alipay.ebpp.bill.add"
}

func (this *AlipayEbppBillAddRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEbppBillAddRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEbppBillAddRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEbppBillAddRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEbppBillAddRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEbppBillAddRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEbppBillAddRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEbppBillAddRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("merchant_order_no", this.MerchantOrderNo)
  txtParams.Put("order_type", this.OrderType)
  txtParams.Put("sub_order_type", this.SubOrderType)
  txtParams.Put("charge_inst", this.ChargeInst)
  txtParams.Put("bill_key", this.BillKey)
  txtParams.Put("owner_name", this.OwnerName)
  txtParams.Put("pay_amount", this.PayAmount)
  txtParams.Put("service_amount", this.ServiceAmount)
  txtParams.Put("bill_date", this.BillDate)
  txtParams.Put("mobile", this.Mobile)
  txtParams.Put("traffic_location", this.TrafficLocation)
  txtParams.Put("traffic_regulations", this.TrafficRegulations)
  txtParams.Put("bank_bill_no", this.BankBillNo)
  txtParams.Put("extend_field", this.ExtendField)
  return txtParams
}
