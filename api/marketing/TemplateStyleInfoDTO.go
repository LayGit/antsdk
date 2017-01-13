package marketing

type TemplateStyleInfoDTO struct {
  CardShowName        string    `json:"card_show_name"`         // 钱包端显示名称（字符串长度）
  LogoId              string    `json:"logo_id"`                // logo的图片ID，通过接口（alipay.offline.material.image.upload）上传图片 图片说明：1M以内，格式bmp、png、jpeg、jpg、gif； 尺寸不小于500*500px的正方形； 请优先使用商家LOGO；
  Color               string    `json:"color"`                  // 卡片颜色
  BackgroundId        string    `json:"background_id"`          // 背景图片Id，通过接口（alipay.offline.material.image.upload）上传图片 图片说明：2M以内，格式：bmp、png、jpeg、jpg、gif； 尺寸不小于1020*643px； 图片不得有圆角，不得拉伸变形
  BgColor             string    `json:"bg_color"`               // 背景色
  FeatureDescriptions []string  `json:"feature_descriptions"`   // 特色信息，用于领卡预览
  Slogan              string    `json:"slogan"`                 // 标语
  SloganImgId         string    `json:"slogan_img_id"`          // 标语图片， 通过接口（alipay.offline.material.image.upload）上传图片
  BrandName           string    `json:"brand_name"`             // 品牌商名称
}
