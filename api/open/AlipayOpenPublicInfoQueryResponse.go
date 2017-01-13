package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicInfoQueryResponse struct {
  api.AlipayResponse
  AppName         string `json:"app_name"`          // 服务窗名称
  LogoUrl         string `json:"logo_url"`          // 服务窗头像地址
  PublicGreeting  string `json:"public_greeting"`   // 服务窗欢迎语
  AuditStatus     string `json:"audit_status"`      // 服务窗审核状态，对于系统商而言，只有四个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过，AUDIT_NORMAL：无审核状态（当前没有处于审核过程的工单）
  AuditDesc       string `json:"audit_desc"`        // 服务窗审核状态描述，如果审核驳回则有相关的驳回理由
  IsOnline        string `json:"is_online"`         // 服务窗是否上线，T表示上线，F表示未上线
  IsRelease       string `json:"is_release"`        // 服务窗是否上架，T表示上架，上架即可在支付宝客户端被搜索到，F表示未上架
}
