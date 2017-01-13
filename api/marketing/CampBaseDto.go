package marketing

type CampBaseDto struct {
  PlanStatus      string              `json:"plan_status"`      // 招商状态,GOING招商中，ENDED招商结束，OFFLINE下架
  AuditStatus     string              `json:"audit_status"`     // 活动审核状态，AUDITING为审核中,REJECT为驳回，不返回为成功
  Id              string              `json:"id"`               // 活动id
  Name            string              `json:"name"`             // 活动名称
  Type            string              `json:"type"`             // 活动类型.DIRECT_SEND:直发奖,CONSUME_SEND:消费送
  AutoDelayFlag   string              `json:"auto_delay_flag"`  // 是否自动续期，Y为是，N为否，为空表示否
  ActivityOrders  []ActivityOrderDTO  `json:"activity_orders"`  // 活动工单列表
  Status          string              `json:"status"`           // 活动状态,CREATED:草稿，ENABLED：生效，DISABLED：无效，STARTED：启动，CLOSED：停止，FINISHED：完成
  StartTime       string              `json:"start_time"`       // 启动时间
  EndTime         string              `json:"end_time"`         // 截止时间
}
