package marketing

type VoucherDescDetail struct {
  Title   string    `json:"title"`    // 券说明的标题
  Details []string  `json:"details"`  // 具体描述信息列表
  Images  []string  `json:"images"`   // 图片描述信息
  Url     string    `json:"url"`      // 券外部详情描述
}
