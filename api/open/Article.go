package open

type Article struct {
  Title       string `json:"title"`       // 图文消息标题
  Desc        string `json:"desc"`        // 图文消息内容
  ImageUrl    string `json:"image_url"`   // 图片链接，对于多条图文消息的第一条消息，该字段不能为空
  Url         string `json:"url"`         // 点击图文消息跳转的链接
  ActionName  string `json:"action_name"` // 链接文字
}
