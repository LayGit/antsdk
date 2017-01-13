package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicInfoModifyResponse struct {
  api.AlipayResponse
  AuditStatus string `json:"audit_status"`  // 服务窗审核状态，对于系统商而言，只有三个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过。
  AuditDesc   string `json:"audit_desc"`    // 服务窗审核状态描述，如果审核驳回则有相关的驳回理由
}
