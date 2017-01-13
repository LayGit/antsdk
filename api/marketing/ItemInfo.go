package marketing

type ItemInfo struct {
  ItemText        string    `json:"item_text"`          // 单品券说明
  ItemName        string    `json:"item_name"`          // 单品名称
  ItemImgs        []string  `json:"item_imgs"`          // 单品图片列表 单品图片不能超过3张
  ItemLink        string    `json:"item_link"`          // 单品券详细介绍跳转链接
  OriginalPrice   string    `json:"original_price"`     // 单品的原价，单位元 必须为合法金额类型字符串，如9.99
  MaxDiscountNum  string    `json:"max_discount_num"`   // 最高优惠商品件数
  MinConsumeNum   string    `json:"min_consume_num"`    // 最低购买商品件数
  ItemIds         []string  `json:"item_ids"`           // 券适用的单品码列表 最少配置1个单品码 最多配置500个单品码
}
