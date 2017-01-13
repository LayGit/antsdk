package marketing

type ActivityOrderDTO struct {
  OrderId             string              `json:"order_id"`             // 订单号
  OrderStatus         string              `json:"order_status"`         // INIT:初始化;DOING:处理中;SUCCESS:成功;FAIL:失败
  AuditStatus         string              `json:"audit_status"`         // INIT:初始化;AUDITING:审核中;REJECT:审核驳回;PASS:审核通过;CANCEL:审核撤销;FAIL:审核失败
  OrderType           string              `json:"order_type"`           // CAMPAIGN_CREATE_ORDER:创建工单;CAMPAIGN_ENABLE_ORDER:生效工单;CAMPAIGN_START_ORDER:启动工单;CAMPAIGN_CLOSE_ORDER:关闭工单;CAMPAIGN_FINISH_ORDER:结束工单;CAMPAIGN_DELETE_ORDER:删除工单;CAMPAIGN_MODIFY_ORDER:修改工单
  ActivityAuditList   []ActivityAuditDTO  `json:"activity_audit_list"`  // 工单中的审核信息
  CreatorId           string              `json:"creator_id"`           // 活动工单创建人
  CreatorType         string              `json:"creator_type"`         // 活动工单创建人类型，PROVIDER:服务商;PROVIDER_STAFF:服务商员工;SALES:BD人员;MER:商户
}
