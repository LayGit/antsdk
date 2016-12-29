package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMarketShopQuerydetailResponse struct {
  api.AlipayResponse
  AuditStatus                 string  `json:"audit_status"`                   // 门店审核状态，对于系统商而言，只有三个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过。第一次审核通过会触发门店上架。
  AuditDesc                   string  `json:"audit_desc"`                     // 门店审核状态描述，如果审核驳回则有相关的驳回理由
  IsOnline                    string  `json:"is_online"`                      // 门店是否上架，T表示上架,F表示未上架，第一次门店审核通过后会触发上架
  IsShow                      string  `json:"is_show"`                        // 门店是否在客户端显示，T表示显示，F表示隐藏
  CategoryId                  string  `json:"category_id"`                    // 类目id，请参考商户入驻要求(https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。
  BrandName                   string  `json:"brand_name"`                     // 品牌名，不填写则默认为“其它品牌”。
  BrandLogo                   string  `json:"brand_logo"`                     // 品牌LOGO; 图片ID，不填写则默认为门店首图main_image。
  MainShopName                string  `json:"main_shop_name"`                 // 主门店名 比如：肯德基；主店名里不要包含分店名，如“万塘路店”。主店名长度不能超过20个字符。
  BranchShopName              string  `json:"main_shop_name"`                 // 分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)。
  ProvinceCode                string  `json:"province_code"`                  // 省份编码，国标码，详见国家统计局数据 点此下载(http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls?spm=a219a.7395905.0.0.SN6mhh&file=2016.xls)。
  CityCode                    string  `json:"city_code"`                      // 城市编码，国标码，详见国家统计局数据 点此下载(http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls?spm=a219a.7395905.0.0.SN6mhh&file=2016.xls)。
  DistrictCode                string  `json:"district_code"`                  // 区县编码，国标码，详见国家统计局数据 点此下载(http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls?spm=a219a.7395905.0.0.SN6mhh&file=2016.xls)。
  Address                     string  `json:"address"`                        // 门店详细地址，地址字符长度在4-50个字符，注：不含省市区。门店详细地址按规范格式填写地址，以免影响门店搜索及活动报名：例1：道路+门牌号，“人民东路18号”；例2：道路+门牌号+标志性建筑+楼层，“四川北路1552号欢乐广场1楼”。
  Longitude                   float64 `json:"longitude,string"`               // 经度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
  Latitude                    float64 `json:"latitude,string"`                // 纬度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
  ContactNumber               string  `json:"contact_number"`                 // 门店电话号码；支持座机和手机，只支持数字和+-号，在客户端对用户展现， 支持多个电话， 以英文逗号分隔。
  StoreId                     string  `json:"store_id"`                       // 外部门店编号；最长32位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性。
  MainImage                   string  `json:"main_image"`                     // 门店首图，非常重要，推荐尺寸2000*1500。
  AuditImages                 string  `json:"audit_images"`                   // 门店审核时需要的图片；至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过；多个图片之间以英文逗号分隔。
  BusinessTime                string  `json:"business_time"`                  // 营业时间，支持分段营业时间，以英文逗号分隔。
  Wifi                        string  `json:"wifi"`                           // 门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示。
  Parking                     string  `json:"parking"`                        // 门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示。
  NoSmoking                   string  `json:"no_smoking"`                     // 是否有无烟区，T表示有无烟区，F表示没有无烟区，不传在客户端不展示。
  Box                         string  `json:"box"`                            // 门店是否有包厢，T表示有，F表示没有，不传在客户端不作展示。
  ValueAdded                  string  `json:"value_added"`                    // 门店其他的服务，门店与用户线下兑现。
  AvgPrice                    string  `json:"avg_price"`                      // 人均消费价格，最少1元，最大不超过99999元，请按实际情况填写；单位元，不需填写单位。
  ISVUid                      string  `json:"isv_uid"`                        // ISV返佣id，门店创建、或者门店交易的返佣将通过此账号反给ISV，如果有口碑签订了返佣协议，则该字段作为返佣数据提取的依据。此字段必须是个合法uid，2088开头的16位支付宝会员账号，如果传入错误将无法创建门店。
  Licence                     string  `json:"licence"`                        // 门店营业执照图片，各行业所需的证照资质参见商户入驻要求(https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。
  LicenceCode                 string  `json:"licence_code"`                   // 门店营业执照编号，营业执照信息与is_operating_online至少填一项。
  LicenceName                 string  `json:"licence_name"`                   // 门店营业执照名称。
  LicenceExpires              string  `json:"licence_expires"`                // 营业执照过期时间。
  AuthLetter                  string  `json:"auth_letter"`                    // 门店授权函,营业执照与签约账号主体不一致时需要。
  IsOperatingOnline           string  `json:"is_operating_online"`            // 是否在其他平台开店，T表示有开店，F表示未开店。
  OnlineImage                 string  `json:"online_image"`                   // 在其他平台的开店图片，支持多张，逗号分隔
  OperateNotifyURL            string  `json:"operate_notify_url"`             // 当商户的门店审核状态发生变化时，会向该地址推送消息。
  ImplementId                 string  `json:"implement_id"`                   // 机具号，多个之间以英文逗号分隔。
  PaymentAccount              string  `json:"payment_account"`                // 门店收款账户，门店收款账户只能被查询，不能通过接口修改。如果为空，则表示门店收款账户为商户签约账户
  QRCode                      string  `json:"qr_code"`                        // 门店收款二维码裸码
  ProcessedQRCode             string  `json:"processed_qr_code"`              // 经过加工后的门店收款二维码
  ProviderXiaoerUid           string  `json:"provider_xiaoer_uid"`            // 门店运营归属人uid
  NotifyMobile                string  `json:"notify_mobile"`                  // 门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息，不在客户端展示；多个以引文逗号分隔
  PicColl                     string  `json:"pic_coll"`                       // 图片集，是map转化成的json串，key是图片id,value是图片url
  BusinessCertificate         string  `json:"business_certificate"`           // 许可证，各行业所需的证照资质参见商户入驻要求；该字段只能上传一张许可证，一张以外的许可证、除营业执照和许可证之外其他证照请放在其他资质字段上传。
  BusinessCertificateExpires  string  `json:"business_certificate_expires"`   // 许可证有效期，格式：2020-03-20。
  PartnerId                   string  `json:"partner_id"`                     // 门店的签约ID
  PayType                     string  `json:"pay_type"`                       // 默认付款类型；如：付款码、扫码付、声波支付、在线买单、其它
  CreateChannel               string  `json:"create_channel"`                 // 门店创建来源；如：开放平台、支付宝客户端、口碑商家app、商家自主开店、服务商开店、全民开店-支付宝客户端、全民开店-商户app、其它
  OtherAuthImages             string  `json:"other_auth_images"`              // 其它资质证明图片集；支持多张，逗号分隔
  GMTShopCreate               string  `json:"gmt_shop_create"`                // 门店创建时间
  GMTShopModified             string  `json:"gmt_shop_modified"`              // 门店修改时间
  ShopTags                    string  `json:"shop_tags"`                      // 门店标签；JSON格式。包括：keyMerchant-是否重点商户（T/F）；isHallMeal-堂食（T/F）；注：若标签 key 不存在，则门店无对应的标签
}
