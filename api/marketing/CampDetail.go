package marketing

type CampDetail struct {
  Type            string              `json:"type"`               // 活动类型.DIRECT_SEND:直发奖,CONSUME_SEND:消费送
  RecruitInfo     RecruitInfo         `json:"recruit_info"`       // 招商信息
  ActivityOrders  []ActivityOrderDTO  `json:"activity_orders"`    // 活动工单列表
  Id              string              `json:"id"`                 // 活动id
  Name            string              `json:"name"`               // 活动名称
  AuditStatus     string              `json:"audit_status"`       // 活动子状态，如审核中
  Status          string              `json:"status"`             // 活动状态,CREATED:草稿，ENABLED：生效，DISABLED：无效，STARTED：启动，CLOSED：停止，FINISHED：完成
  StartTime       string              `json:"start_time"`         // 活动开始时间
  AutoDelayFlag   string              `json:"auto_delay_flag"`    // 是否自动续期该活动,Y表示是，N表示否，默认为N
  EndTime         string              `json:"end_time"`           // 活动结束时间
  Desc            string              `json:"desc"`               // 活动描述
  BudgetInfo      BudgetInfo          `json:"budget_info"`        // 预算信息
  ConstraintInfo  ConstraintInfo      `json:"constraint_info"`    // 活动约束信息
  PromoTools      []PromoTool         `json:"promo_tools"`        // 营销工具
  PublishChannels []PublishChannel    `json:"publish_channels"`   // 投放渠道信息
  ExtInfo         string              `json:"ext_info"`           // 扩展参数
}
