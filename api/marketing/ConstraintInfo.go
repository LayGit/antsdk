package marketing

type ConstraintInfo struct {
  UserWinCount      string    `json:"user_win_count"`       // 活动期间用户能够参与的次数限制 如果不设置则不限制参与次数
  UserWinFrequency  string    `json:"user_win_frequency"`   // 活动期间用户能够参与的频率限制 如果不设置则不限制参与频率 每日中奖1次: D||1 每周中奖2次: W||2 每月中奖3次: M||3
  CrowdGroupId      string    `json:"crowd_group_id"`       // 人群规则组ID 仅直发奖类型活动设置有效，通过调用营销活动人群组规则创建接口参数返回
  CrowdRestriction  string    `json:"crowd_restriction"`    // 针对指定人群的约束条件
  SuitShops         []string  `json:"suit_shops"`           // 活动适用的门店列表 仅品牌商发起的招商活动可为空 最多支持10w家门店
  MinCost           string    `json:"min_cost"`             // 最低消费金额，单位元 仅在创建消费送礼包活动时设置
  ItemIds           []string  `json:"item_ids"`             // 单品码列表 仅在创建消费单品送活动时设置，最多设置500个单品码,由商户根据自己的商品管理自定义，一般为国标码
}
