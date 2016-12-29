package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMarketShopModifyResponse struct {
  api.AlipayResponse
  AuditStatus string `json:"audit_status"`  // 同步请求如果支付宝受理成功，将返回AUDITING状态。
  ApplyId     string `json:"apply_id"`      // 修改门店请求受理成功后返回的支付宝流水ID，根据此ID调用接口 alipay.offline.market.applyorder.batchquery，能够获取当前修改店铺请求审核状态、驳回原因等信息。
}
