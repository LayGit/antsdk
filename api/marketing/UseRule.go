package marketing

type UseRule struct {
  UseTime       []UseTime     `json:"use_time"`       // 券可用时间段
  ForbiddenTime ForbiddenTime `json:"forbidden_time"` // 券的不可用时间
  SuitShops     []string      `json:"suit_shops"`     // 券适用门店列表 仅品牌商发起的招商活动可为空 直发奖类型活动必须与活动适用门店一致 最多支持10w家门店
  MinConsume    string        `json:"min_consume"`    // 券核销的最低消费门槛，单位元
  ExtInfo       string        `json:"ext_info"`       // 扩展属性，无需设置
}

type UseTime struct {
  Dimension string `json:"dimension"` // 券可用时段时间维度，目前支持周(W)
  Values    string `json:"values"`    // 券可用时间维度值 周维度的取值范围1-7(周一至周日)，多个可用时段用逗号分隔 如"1,3,5"，对应周一，周三，周五可用
  Times     string `json:"times"`     // 券可用时间段 可用时间段起止时间用逗号分隔，多个时间段之间用^分隔 如, "16:00:00,20:00:00^21:00:00,22:00:00"表示16点至20点，21点至22点可用 时间段不可重叠
}

type ForbiddenTime struct {
  Days string `json:"days"` // 不可用日期区间，仅支持到天 不可用区间起止日期用逗号分隔，多个区间之间用^分隔 如"2016-05-01,2016-05-03^2016-10-01,2016-10-07"表示2016年5月1日至5月3日，10月1日至10月7日券不可用
}
