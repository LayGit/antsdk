package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppBillAddResponse struct {
  api.AlipayResponse
  MerchantOrderNo     string  `json:"merchant_order_no"`        // 输出机构的业务流水号，需要保证唯一性。
  AlipayOrderNo       string  `json:"alipay_order_no"`          // 支付宝的业务订单号，具有唯一性。
  OrderType           string  `json:"order_type"`               // 支付宝订单类型。公共事业缴纳JF,信用卡还款HK
  SubOrderType        string  `json:"sub_order_type"`           // 子业务类型是业务类型的下一级概念，例如：WATER表示JF下面的水费，ELECTRIC表示JF下面的电费，GAS表示JF下面的燃气费。
  ChargeInst          string  `json:"charge_inst"`              // 支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
  ChargeInstName      string  `json:"charge_inst_name"`         // 出账机构中文名称。
  BillKey             string  `json:"bill_key"`                 // 账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
  OwnerName           string  `json:"owner_name"`               // 拥有该账单的用户姓名
  PayAmount           float64 `json:"pay_amount,string"`        // 缴费金额。用户支付的总金额。单位为：RMB Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。
  ServiceAmount       float64 `json:"service_amount,string"`    // 账单的服务费
  BillDate            string  `json:"bill_date"`                // 账单的账期，例如201203表示2012年3月的账单。
  BankBillNo          string  `json:"bank_bill_no"`             // 外部订单号，由于对账时回传给外部商户
  ExtendField         string  `json:"extend_field"`             // 扩展属性，该属性值现在用于确保只有一个人可以支付成功 用法：多个人对同一笔外部欠费单创建多个账单时，确保该值不变
}
