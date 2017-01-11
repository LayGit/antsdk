package car

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 车辆驶入接口
// 上传车辆驶入信息，上传信息通过该接口提交到支付宝，支付宝返回对应的信息
type AlipayEcoMycarParkingEnterinfoSyncRequest struct {
  api.IAlipayRequest
  TerminalType        string                                               `json:"terminal_type"`
  TerminalInfo        string                                               `json:"terminal_info"`
  ProdCode            string                                               `json:"prod_code"`
  NotifyUrl           string                                               `json:"notify_url"`
  ReturnUrl           string                                               `json:"return_url"`
  BizContent          AlipayEcoMycarParkingEnterinfoSyncRequestBizContent  `json:"biz_content"`
}

type AlipayEcoMycarParkingEnterinfoSyncRequestBizContent struct {
  ParkingId string `json:"parking_id"`  // 支付宝停车场ID ，系统唯一
  CarNumber string `json:"car_number"`  // 车牌号
  InTime    string `json:"in_time"`     // 车辆入场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetApiMethodName() string {
  return "alipay.eco.mycar.parking.enterinfo.sync"
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEcoMycarParkingEnterinfoSyncRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
