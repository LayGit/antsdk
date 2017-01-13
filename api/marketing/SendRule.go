package marketing

type SendRule struct {
  MinCost         string `json:"min_cost"`          // 发券最低消费金额，单位元 活动类型为消费送且不是消费送礼包时设置 多营销工具之间不允许设置重复值
  SendNum         string `json:"send_num"`          // 发券数目 最少发1张券，最多发5张券
  SendBudget      string `json:"send_budget"`       // 券的预算数量（仅对口令送随机抽奖有效，即当活动类型为GUESS_SEND，且营销工具PromoTool的个数大于1时，此字段必填，其余情况此字段必为空）
  AllowRepeatSend string `json:"allow_repeat_send"` // 是否允许重复发奖： true代表允许，false代表不允许 默认不设置，表明用户领取券后如果没有核销则不允许再次领取券 如果设置为true，表明如果用户领取券后没有核销，还可以继续领取该券
}
