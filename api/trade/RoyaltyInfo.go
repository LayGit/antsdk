package trade

type RoyaltyInfo struct {
  RoyaltyType         string                `json:"royalty_type"`           // 分账类型 卖家的分账类型，目前只支持传入ROYALTY（普通分账类型）。
  RoyaltyDetailInfos  []RoyaltyDetailInfos  `json:"royalty_detail_infos"`   // 分账明细的信息，可以描述多条分账指令
}

type RoyaltyDetailInfos struct {
  SerialNo          int     `json:"serial_no,string"`          // 分账序列号，表示分账执行的顺序，必须为正整数
  TransInType       string  `json:"trans_in_type"`             // 接受分账金额的账户类型： 	userId：支付宝账号对应的支付宝唯一用户号。 	bankIndex：分账到银行账户的银行编号。目前暂时只支持分账到一个银行编号。 storeId：分账到门店对应的银行卡编号。 默认值为userId。
  BatchNo           string  `json:"batch_no"`                  // 分账批次号 分账批次号。 目前需要和转入账号类型为bankIndex配合使用。
  OutRelationId     string  `json:"out_relation_id"`           // 商户分账的外部关联号，用于关联到每一笔分账信息，商户需保证其唯一性。 如果为空，该值则默认为“商户网站唯一订单号+分账序列号”
  TransOutType      string  `json:"trans_out_type"`            // 要分账的账户类型。 目前只支持userId：支付宝账号对应的支付宝唯一用户号。 默认值为userId。
  TransOut          string  `json:"trans_out"`                 // 如果转出账号类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
  TransIn           string  `json:"trans_in"`                  // 如果转入账号类型为userId，本参数为接受分账金额的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。 	如果转入账号类型为bankIndex，本参数为28位的银行编号（商户和支付宝签约时确定）。 如果转入账号类型为storeId，本参数为商户的门店ID。
  Amount            float64 `json:"amount,string"`             // 分账的金额，单位为元
  Desc              string  `json:"desc"`                      // 分账描述信息
  AmountPercentage  int     `json:"amount_percentage,string"`  // 分账的比例，值为20代表按20%的比例分账
}
