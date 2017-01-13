package marketing

type ActivityAuditDTO struct {
  AuditId       string `json:"audit_id"`        // 审核id
  AuditStatus   string `json:"audit_status"`    // INIT:初始化;AUDITING:审核中;REJECT:审核驳回;PASS:审核通过;CANCEL:审核撤销;FAIL:审核失败
  Reason        string `json:"reason"`          // 审核通过或者审核驳回的原因
  OperationTime string `json:"operation_time"`  // 操作时间
  CreatorId     string `json:"creator_id"`      // 操作人id
  CreatorType   string `json:"creator_type"`    // SALES:口碑内部小二;PROVIDER:外部服务商;PROVIDER_STAFF:外部服务商员工
}
