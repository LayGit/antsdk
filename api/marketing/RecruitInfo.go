package marketing

type RecruitInfo struct {
  PlanId                  string  `json:"plan_id"`                    // 招商方案id
  ExcludeConstraintShops  bool    `json:"exclude_constraint_shops"`   // 是否参与门店参与了招商
  StartTime               string  `json:"start_time"`                 // 招商开始时间
  EndTime                 string  `json:"end_time"`                   // 招商结束时间
  Status                  string  `json:"status"`                     // 招商状态
}
