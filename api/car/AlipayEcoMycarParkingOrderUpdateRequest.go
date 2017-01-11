package car

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 订单更新接口
// 商户通过接口调用，回传订单状态给停车平台
type AlipayEcoMycarParkingOrderUpdateRequest struct {
  api.IAlipayRequest
  TerminalType        string                                            `json:"terminal_type"`
  TerminalInfo        string                                            `json:"terminal_info"`
  ProdCode            string                                            `json:"prod_code"`
  NotifyUrl           string                                            `json:"notify_url"`
  ReturnUrl           string                                            `json:"return_url"`
  BizContent          AlipayEcoMycarParkingOrderUpdateRequestBizContent `json:"biz_content"`
}

type AlipayEcoMycarParkingOrderUpdateRequestBizContent struct {
  UserId      string `json:"user_id"`       // 停车缴费支付宝用户的ID，请ISV保证用户ID的正确性，以免导致用户在停车平台查询不到相关的订单信息
  OrderNo     string `json:"order_no"`      // 支付宝支付流水号，系统唯一
  OrderStatus string `json:"order_status"`  // 用户停车订单状态，0：成功，1：失败
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetApiMethodName() string {
  return "alipay.eco.mycar.parking.order.update"
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEcoMycarParkingOrderUpdateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
