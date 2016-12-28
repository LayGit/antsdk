package trade

type OpenApiRoyaltyDetailInfoPojo struct {
  TransOut          string  `json:"trans_out"`                 // 分账支出方账户，类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
  TransIn           string  `json:"trans_in"`                  // 分账收入方账户，类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
  Amount            float64 `json:"amount,string"`             // 分账的金额，单位为元
  AmountPercentage  int     `json:"amount_percentage,string"`  // 分账信息中分账百分比。取值范围为大于0，少于或等于100的整数。
  Desc              string  `json:"desc"`                      // 分账描述
}
