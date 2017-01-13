package marketing

type PromoTool struct {
  Voucher   Voucher   `json:"voucher"`      // 券对象，当活动类型为POINT_SEND时为null，其他活动类型此字段必填
  Status    string    `json:"status"`       // 单个营销工具的生效状态，当在招商部分券失效后会使用这个字段
  VoucherNo string    `json:"voucher_no"`   // 单个营销工具的生效状态，当在招商部分券失效后会使用这个字段
  PointCard PointCard `json:"point_card"`   // 集点卡工具，仅在活动类型为POINT_SEND时才有效且必填，其他活动类型此字段必须为null
  SendRule  SendRule  `json:"send_rule"`    // 奖品发放的规则
}
