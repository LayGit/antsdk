package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 业务流水批量查询接口
// 通过该接口分页查询Leads、门店、商品相关操作流水信息
type AlipayOfflineMarketApplyorderBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                                                    `json:"terminal_type"`
  TerminalInfo  string                                                    `json:"terminal_info"`
  ProdCode      string                                                    `json:"prod_code"`
  NotifyUrl     string                                                    `json:"notify_url"`
  ReturnUrl     string                                                    `json:"return_url"`
  BizContent    AlipayOfflineMarketApplyorderBatchqueryRequestBizContent  `json:"biz_content"`
}

type AlipayOfflineMarketApplyorderBatchqueryRequestBizContent struct {
  ApplyIds    []string  `json:"apply_ids"`      // 支付宝流水ID列表，最大不超过100个
  RequestIds  []string  `json:"request_ids"`    // 请求ID列表，最大不超过100个。 注意：暂时不支持此字段查询。
  BizId       string    `json:"biz_id"`         // 业务主体ID。根据biz_type不同可能对应shop_id或item_id。
  BizType     string    `json:"biz_type"`       // 业务类型：SHOP-店铺，ITEM-商品。
  Action      string    `json:"action"`         // 操作动作
  OPId        string    `json:"op_id"`          // 操作用户的支付账号id
  Status      string    `json:"status"`         // 流水状态：INIT-初始，PROCESS-处理中，SUCCESS-成功，FAIL-失败。
  StartTime   string    `json:"start_time"`     // 查询的流水创建时间起始值，只能查询近3个月数据。格式：yyyy-MM-dd HH:mm:ss
  EndTime     string    `json:"end_time"`       // 查询的流水创建时间最后值。格式：yyyy-MM-dd HH:mm:ss
  OPRole      string    `json:"op_role"`        // 系统集成商统一传入ISV
  PageNo      int       `json:"page_no"`        // 页码，留空标示第一页，默认20个结果为一页
  PageSize    int       `json:"page_size"`      // 每页记录数。默认20，最大100。
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetApiMethodName() string {
  return "alipay.offline.market.applyorder.batchquery"
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMarketApplyorderBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
