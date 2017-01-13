package marketing

type RecruitTool struct {
  StartTime               string        `json:"start_time"`
  ExcludeConstraintShops  bool          `json:"exclude_constraint_shops"`
  EndTime                 string        `json:"end_time"`
  PidShops                []PidShopInfo `json:"pid_shops"`
}

type PidShopInfo struct {
  Pid     string    `json:"pid"`        // 商户pid
  ShopIds []string  `json:"shop_ids"`   // pid下的门店列表
}
