package ebpp

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEbppBillGetResponse struct {
  api.AlipayResponse
  MerchantOrderNo     string  `json:"merchant_order_no"`        // 输出机构的业务流水号，需要保证唯一性。
  AlipayOrderNo       string  `json:"alipay_order_no"`          // 支付宝的业务订单号，具有唯一性。
  OrderStatus         string  `json:"order_status"`             // 账单的状态。 INIT:等待付款，SUCCESS:成功,FAILED:失败。
  OrderType           string  `json:"order_type"`               // 支付宝订单类型。公共事业缴纳JF,信用卡还款HK
  SubOrderType        string  `json:"sub_order_type"`           // 子业务类型是业务类型的下一级概念，例如：WATER表示JF下面的水费，ELECTRIC表示JF下面的电费，GAS表示JF下面的燃气费。
  ChargeInst          string  `json:"charge_inst"`              // 支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
  ChargeInstName      string  `json:"charge_inst_name"`         // 出账机构中文名称。
  BillKey             string  `json:"bill_key"`                 // 账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
  OwnerName           string  `json:"owner_name"`               // 拥有该账单的用户姓名
  PayAmount           float64 `json:"pay_amount,string"`        // 缴费金额。用户支付的总金额。单位为：RMB Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。
  ServiceAmount       float64 `json:"service_amount,string"`    // 账单的服务费
  BillDate            string  `json:"bill_date"`                // 账单的账期，例如201203表示2012年3月的账单。
  TrafficLocation     string  `json:"traffic_location"`         // 交通违章地点，sub_order_type=TRAFFIC时有值
  TrafficRegulations  string  `json:"traffic_regulations"`      // 违章行为，sub_order_type=TRAFFIC时有值。
  PayTime             string  `json:"pay_time"`                 // 付款时间
}
