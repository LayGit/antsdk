package marketing

type MerchantCard struct {
  BizCardNo       string `json:"biz_card_no"`       // 支付宝业务卡号 说明： 1、开卡成功后返回该参数，需要保存留用； 2、开卡/更新/删卡/查询卡接口请求中不需要传该参数；
  ExternalCardNo  string `json:"external_card_no"`  // 商户外部会员卡卡号 说明： 1、会员卡开卡接口，如果卡类型为外部会员卡，请求中则必须提供该参数； 2、更新、查询、删除等接口，请求中则不需要提供该参数值；
  OpenDate        string `json:"open_date"`         // 会员卡开卡时间，格式为yyyy-MM-dd HH:mm:ss
  ValidDate       string `json:"valid_date"`        // 会员卡有效期
  Level           string `json:"level"`             // 会员卡等级（由商户自定义，并可以在卡模板创建时，定义等级信息）
  Point           string `json:"point"`             // 会员卡积分，积分必须为数字型（可为浮点型，带2位小数点）
  Balance         string `json:"balance"`           // 资金卡余额，单位：元，精确到小数点后两位。
}
