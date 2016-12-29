package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMarketApplyorderBatchqueryResponse struct {
  api.AlipayResponse
  PageSize      int                     `json:"page_size,string"`       // 每页记录数
  CurrentPageNo int                     `json:"total_page_no,string"`   // 当前页码
  TotalPageNo   int                     `json:"total_page_no,string"`   // 总页码数目
  TotalItems    int                     `json:"total_items,string"`     // 总记录数
  BizOrderInfos []BizOrderQueryResponse `json:"biz_order_infos"`        // 支付宝操作流水信息列表
}

type BizOrderQueryResponse struct {
  ApplyId         string `json:"apply_id"`          // 支付宝流水ID
  RequestId       string `json:"request_id"`        // 注意：此字段并非外部商户请求时传入的request_id，暂时代表支付宝内部字段，请勿用。
  BizType         string `json:"biz_type"`          // 业务类型：SHOP-店铺，ITEM-商品
  BizId           string `json:"biz_id"`            // 业务主体ID。根据biz_type不同可能对应shop_id或item_id。 特别注意对于门店创建，当流水status=SUCCESS时，此字段才为shop_id，其他状态时为0或空。
  Action          string `json:"action"`            // 操作动作。 CREATE_SHOP-创建门店， MODIFY_SHOP-修改门店， CREATE_ITEM-创建商品， MODIFY_ITEM-修改商品， EFFECTIVE_ITEM-上架商品， INVALID_ITEM-下架商品， RESUME_ITEM-暂停售卖商品， PAUSE_ITEM-恢复售卖商品
  ActionMode      string `json:"action_mode"`       // 操作模式：NORMAL-普通开店
  Status          string `json:"status"`            // 流水状态：INIT-初始，PROCESS-处理中，SUCCESS-成功，FAIL-失败。
  SubStatus       string `json:"sub_status"`        // 流水子状态：WAIT_CERTIFY-等待认证，LICENSE_AUDITING-证照审核中，RISK_AUDITING-风控审核中，WAIT_SIGN-等待签约，FINISH-终结。
  OPId            string `json:"op_id"`             // 操作用户的支付账号id
  ResultCode      string `json:"result_code"`       // 流水处理结果码 点此查看(https://doc.open.alipay.com/doc2/detail.htm?spm=a219a.7629140.0.0.lL9hGI&treeId=78&articleId=103834&docType=1#s2)
  ResultDesc      string `json:"result_desc"`       // 流水处理结果描述
  CreateTime      string `json:"create_time"`       // 创建时间
  UpdateTime      string `json:"update_time"`       // 更新时间
  BizContextInfo  string `json:"biz_context_info"`  // 流水上下文信息，JSON格式。根据action不同对应的结构也不同，JSON字段与含义可参考各个接口的请求参数。
}
