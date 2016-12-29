package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 创建门店信息
// 系统商需要通过该接口在口碑平台帮助商户创建门店信息。
type AlipayOfflineMarketShopCreateRequest struct {
  api.IAlipayRequest
  TerminalType  string                                          `json:"terminal_type"`
  TerminalInfo  string                                          `json:"terminal_info"`
  ProdCode      string                                          `json:"prod_code"`
  NotifyUrl     string                                          `json:"notify_url"`
  ReturnUrl     string                                          `json:"return_url"`
  BizContent    AlipayOfflineMarketShopCreateRequestBizContent  `json:"biz_content"`
}

type AlipayOfflineMarketShopCreateRequestBizContent struct {
  StoreId                     string  `json:"store_id"`                       // 外部门店编号；最长32位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性。
  CategoryId                  string  `json:"category_id"`                    // 类目id，请参考商户入驻要求(https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。
  BrandName                   string  `json:"brand_name"`                     // 品牌名，不填写则默认为“其它品牌”。
  BrandLogo                   string  `json:"brand_logo"`                     // 品牌LOGO; 图片ID，不填写则默认为门店首图main_image。
  MainShopName                string  `json:"main_shop_name"`                 // 主门店名 比如：肯德基；主店名里不要包含分店名，如“万塘路店”。主店名长度不能超过20个字符。
  BranchShopName              string  `json:"main_shop_name"`                 // 分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)。
  ProvinceCode                string  `json:"province_code"`                  // 省份编码，国标码，详见国家统计局数据 点此下载(http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls?spm=a219a.7395905.0.0.SN6mhh&file=2016.xls)。
  CityCode                    string  `json:"city_code"`                      // 城市编码，国标码，详见国家统计局数据 点此下载(http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls?spm=a219a.7395905.0.0.SN6mhh&file=2016.xls)。
  DistrictCode                string  `json:"district_code"`                  // 区县编码，国标码，详见国家统计局数据 点此下载(http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls?spm=a219a.7395905.0.0.SN6mhh&file=2016.xls)。
  Address                     string  `json:"address"`                        // 门店详细地址，地址字符长度在4-50个字符，注：不含省市区。门店详细地址按规范格式填写地址，以免影响门店搜索及活动报名：例1：道路+门牌号，“人民东路18号”；例2：道路+门牌号+标志性建筑+楼层，“四川北路1552号欢乐广场1楼”。
  Longitude                   float64 `json:"longitude"`                      // 经度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
  Latitude                    float64 `json:"latitude"`                       // 纬度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
  ContactNumber               string  `json:"contact_number"`                 // 门店电话号码；支持座机和手机，只支持数字和+-号，在客户端对用户展现， 支持多个电话， 以英文逗号分隔。
  NotifyMobile                string  `json:"notify_mobile"`                  // 门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息， 不在客户端展示。
  MainImage                   string  `json:"main_image"`                     // 门店首图，非常重要，推荐尺寸2000*1500。
  AuditImages                 string  `json:"audit_images"`                   // 门店审核时需要的图片；至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过；多个图片之间以英文逗号分隔。
  BusinessTime                string  `json:"business_time"`                  // 营业时间，支持分段营业时间，以英文逗号分隔。
  Wifi                        string  `json:"wifi"`                           // 门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示。
  Parking                     string  `json:"parking"`                        // 门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示。
  ValueAdded                  string  `json:"value_added"`                    // 门店其他的服务，门店与用户线下兑现。
  AvgPrice                    string  `json:"avg_price"`                      // 人均消费价格，最少1元，最大不超过99999元，请按实际情况填写；单位元，不需填写单位。
  ISVUid                      string  `json:"isv_uid"`                        // ISV返佣id，门店创建、或者门店交易的返佣将通过此账号反给ISV，如果有口碑签订了返佣协议，则该字段作为返佣数据提取的依据。此字段必须是个合法uid，2088开头的16位支付宝会员账号，如果传入错误将无法创建门店。
  Licence                     string  `json:"licence"`                        // 门店营业执照图片，各行业所需的证照资质参见商户入驻要求(https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。
  LicenceCode                 string  `json:"licence_code"`                   // 门店营业执照编号，营业执照信息与is_operating_online至少填一项。
  LicenceName                 string  `json:"licence_name"`                   // 门店营业执照名称。
  LicenceExpires              string  `json:"licence_expires"`                // 营业执照过期时间。
  BusinessCertificate         string  `json:"business_certificate"`           // 许可证，各行业所需的证照资质参见商户入驻要求；该字段只能上传一张许可证，一张以外的许可证、除营业执照和许可证之外其他证照请放在其他资质字段上传。
  BusinessCertificateExpires  string  `json:"business_certificate_expires"`   // 许可证有效期，格式：2020-03-20。
  AuthLetter                  string  `json:"auth_letter"`                    // 门店授权函,营业执照与签约账号主体不一致时需要。
  IsOperatingOnline           string  `json:"is_operating_online"`            // 是否在其他平台开店，T表示有开店，F表示未开店。
  OnlineURL                   string  `json:"online_url"`                     // 其他平台开店的店铺链接url，多个url使用英文逗号隔开,isv迁移到新接口使用此字段，与is_operating_online=T配套使用。
  OperateNotifyURL            string  `json:"operate_notify_url"`             // 当商户的门店审核状态发生变化时，会向该地址推送消息。
  ImplementId                 string  `json:"implement_id"`                   // 机具号，多个之间以英文逗号分隔。
  NoSmoking                   string  `json:"no_smoking"`                     // 是否有无烟区，T表示有无烟区，F表示没有无烟区，不传在客户端不展示。
  Box                         string  `json:"box"`                            // 门店是否有包厢，T表示有，F表示没有，不传在客户端不作展示。
  RequestId                   string  `json:"request_id"`                     // 支持英文字母和数字，由开发者自行定义（不允许重复），在门店notify消息中也会带有该参数，以此标明本次notify消息是对哪个请求的回应。
  OtherAuthorization          string  `json:"other_authorization"`            // 其他资质。用于上传营业证照、许可证照外的其他资质，除已上传许可证外的其他许可证也可以在该字段上传。
  OPRole                      string  `json:"op_role"`                        // 表示以系统集成商的身份开店，开放平台现在统一传入ISV。
  BizVersion                  string  `json:"biz_version"`                    // 店铺接口业务版本号，新接入的ISV，请统一传入2.0。
}

func (this *AlipayOfflineMarketShopCreateRequest) GetApiMethodName() string {
  return "alipay.offline.market.shop.create"
}

func (this *AlipayOfflineMarketShopCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMarketShopCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMarketShopCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMarketShopCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMarketShopCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMarketShopCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMarketShopCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
