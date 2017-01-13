package marketing

type TemplateCardLevelConfDTO struct {
  Level         string `json:"level"`             // 会员级别 该级别和开卡接口中的levle要一致
  LevelShowName string `json:"level_show_name"`   // 会员级别显示名称
  LevelIcon     string `json:"level_icon"`        // 会员级别对应icon， 通过接口（alipay.offline.material.image.upload）上传图片
  LevelDesc     string `json:"level_desc"`        // 会员级别描述
}
