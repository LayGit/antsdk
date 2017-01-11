package car

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 订单同步接口
// 商户通过接口调用，回传订单信息给停车平台
type AlipayEcoMycarParkingOrderSyncRequest struct {
  api.IAlipayRequest
  TerminalType        string                                          `json:"terminal_type"`
  TerminalInfo        string                                          `json:"terminal_info"`
  ProdCode            string                                          `json:"prod_code"`
  NotifyUrl           string                                          `json:"notify_url"`
  ReturnUrl           string                                          `json:"return_url"`
  BizContent          AlipayEcoMycarParkingOrderSyncRequestBizContent `json:"biz_content"`
}

type AlipayEcoMycarParkingOrderSyncRequestBizContent struct {
  UserId        string `json:"user_id"`           // 停车缴费支付宝用户的ID，请ISV保证用户ID的正确性，以免导致用户在停车平台查询不到相关的订单信息
  OutParkingId  string `json:"out_parking_id"`    // ISV停车场ID，由ISV提供，同一个isv或商户范围内唯一
  ParkingName   string `json:"parking_name"`      // 停车场名称，由ISV定义，尽量与高德地图上的一致
  CarNumber     string `json:"car_number"`        // 车牌
  OutOrderNo    string `json:"out_order_no"`      // 设备商订单号，由ISV系统生成
  OrderStatus   string `json:"order_status"`      // 设备商订单状态，0：成功，1：失败
  OrderTime     string `json:"order_time"`        // 订单创建时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
  OrderNo       string `json:"order_no"`          // 支付宝支付流水，系统唯一
  PayTime       string `json:"pay_time"`          // 缴费时间, 格式"YYYYMM-DD HH:mm:ss"，24小时制
  PayType       string `json:"pay_type"`          // 付款方式，1：支付宝在线缴费 ，2：支付宝代扣缴费
  PayMoney      string `json:"pay_money"`         // 缴费金额，保留小数点后两位
  InTime        string `json:"in_time"`           // 入场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
  ParkingId     string `json:"parking_id"`        // 支付宝停车场id，系统唯一
  InDuration    string `json:"in_duration"`       // 停车时长（以分为单位）
  CardNumber    string `json:"card_number"`       // 如果是停车卡缴费，则填入停车卡卡号，否则为'*'
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetApiMethodName() string {
  return "alipay.eco.mycar.parking.order.sync"
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEcoMycarParkingOrderSyncRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
