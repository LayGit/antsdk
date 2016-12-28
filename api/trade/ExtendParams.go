package trade

type ExtendParams struct {
  SysServiceProviderId  string `json:"sys_service_provider_id"`   // 系统商编号 该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
  HBFQNum               string `json:"hb_fq_num"`                 // 使用花呗分期要进行的分期数
  HBFQSellerPercent     string `json:"hb_fq_seller_percent"`      // 使用花呗分期需要卖家承担的手续费比例的百分值，传入100代表100%
}
