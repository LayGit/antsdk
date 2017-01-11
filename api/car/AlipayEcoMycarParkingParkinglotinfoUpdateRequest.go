package car

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

type AlipayEcoMycarParkingParkinglotinfoUpdateRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                      `json:"terminal_type"`
  TerminalInfo        string                                                      `json:"terminal_info"`
  ProdCode            string                                                      `json:"prod_code"`
  NotifyUrl           string                                                      `json:"notify_url"`
  ReturnUrl           string                                                      `json:"return_url"`
  BizContent          AlipayEcoMycarParkingParkinglotinfoUpdateRequestBizContent  `json:"biz_content"`
}

type AlipayEcoMycarParkingParkinglotinfoUpdateRequestBizContent struct {
  ParkingId             string `json:"parking_id"`                // 支付宝返回停车场id，系统唯一
  CityId                string `json:"city_id"`                   // 城市编号（国家统一标准编码）
  EquipmentName         string `json:"equipment_name"`            // 设备商名称
  OutParkingId          string `json:"out_parking_id"`            // ISV停车场ID，由ISV提供，同一个ISV或商户范围内唯一
  ParkingAddress        string `json:"parking_address"`           // 停车场地址
  Longitude             string `json:"longitude"`                 // 经度，最长15位字符(包括小数点)，注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数的准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
  Latitude              string `json:"latitude"`                  // 纬度，最长15位字符(包括小数点)，注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数的准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
  ParkingStartTime      string `json:"parking_start_time"`        // 停车场开始营业时间，格式 "HH:mm:ss"
  ParkingEndTime        string `json:"parking_end_time"`          // 停车场结束营业时间，格式 "HH:mm:ss"
  ParkingName           string `json:"parking_name"`              // 停车场名称，由ISV定义，尽量与高德地图上的一致
  ParkingNumber         string `json:"parking_number"`            // 停车位数目
  ParkingLotType        string `json:"parking_lot_type"`          // 停车场类型，1为小区停车场、2为商圈停车场、3为路面停车场、4为园区停车场、5为写字楼停车场、6为私人停车场
  ParkingType           string `json:"parking_type"`              // 停车场类型(1为地面，2为地下，3为路边)（多个类型，中间用,隔开
  PaymentMode           string `json:"payment_mode"`              // 缴费模式（1为停车卡缴费，2为物料缴费，3为中央缴费机）
  PayType               string `json:"pay_type"`                  // 支付方式（1为支付宝在线缴费，2为支付宝代扣缴费，3当面付)，如支持多种方式以','进行间隔
  ShopingmallId         string `json:"shopingmall_id"`            // 商圈id
  ParkingFeeDescription string `json:"parking_fee_description"`   // 收费说明
  ContactName           string `json:"contact_name"`              // 停车场联系人，如果有则填入
  ContactMobile         string `json:"contact_mobile"`            // 停车场联系人手机，如果有则填入
  ContactTel            string `json:"contact_tel"`               // 停车场联系人座机，如果有则填入
  ContactEmail          string `json:"contact_email"`             // 停车场联系人邮箱，如果有则填入
  ContactWeixin         string `json:"contact_weixin"`            // 停车场联系人微信，如果有则填入
  ContactAlipay         string `json:"contact_alipay"`            // 停车场联系人支付宝账户，如果有则填入
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetApiMethodName() string {
  return "alipay.eco.mycar.parking.parkinglotinfo.update"
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEcoMycarParkingParkinglotinfoUpdateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
