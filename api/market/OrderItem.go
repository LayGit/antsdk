package market

type OrderItem struct {
  Creator       string `json:"creator"`         // 门店创建人(已删除)
  PhoneNO       string `json:"phone_no"`        // 订单所属人联系方式（手机或者座机）
  OnlineTime    string `json:"online_time"`     // 上架时间
  ExpireDate    string `json:"expire_date"`     // 过期时间
  OrderStatus   string `json:"order_status"`    // TO_DO-未实施,DOING-实施中,TO_CONFIRM-待商户确认,DONE-已完成,MERCHANT_REJECTED-商户已回绝,MERCHANT_CANCELLED-商户已取消,ISV_REJECTED-服务商已回绝,ISV_CANCELLED-服务商已取消
  ShopStatus    string `json:"shop_status"`     // 店铺状态（ONLINE--已上架 OFFLINE--未上架 AVAILABLE--已开通 INIT--未开通 EXPIRED--已过期）
  Status        string `json:"status"`          // 待服务商接单
  CommodityId   string `json:"commodity_id"`    // 订购的服务商品ID
  MerchantPid   string `json:"merchant_pid"`    // 商户PID
  MerchantName  string `json:"merchant_name"`   // 商户名称
  BrandName     string `json:"brand_name"`      // 品牌名称
  Contacts      string `json:"contacts"`        // 订单联系人
  ShopName      string `json:"shop_name"`       // 店铺名称
  ShopId        string `json:"shop_id"`         // 店铺ID
  Category      string `json:"category"`        // 店铺品类
  Province      string `json:"province"`        // 店铺所在的省份
  City          string `json:"city"`            // 店铺所在的市
  Address       string `json:"address"`         // 店铺所在具体位置
}
