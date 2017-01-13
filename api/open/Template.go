package open

type Template struct {
  TemplateId  string `json:"template_id"` // 消息模板ID
  Context     Context `json:"context"`    // 消息模板上下文，即模板中定义的参数及参数值
}

type Context struct {
  HeadColor   string  `json:"head_color"`    // 顶部色条的色值
  Url         string  `json:"url"`           // 点击消息后承接页的地址
  ActionName  string  `json:"action_name"`   // 底部链接描述文字，如“查看详情”
  Keyword1    Keyword `json:"keyword1"`      // 模板中占位符的值及文字颜色，value和color都为必填项，color为当前文字颜色
  Keyword2    Keyword `json:"keyword2"`      // 模板中占位符的值及文字颜色，value和color都为必填项，color为当前文字颜色
  First       Keyword `json:"first"`         // 模板中占位符的值及文字颜色，value和color都为必填项，color为当前文字颜色
  Remark      Keyword `json:"remark"`         // 模板中占位符的值及文字颜色，value和color都为必填项，color为当前文字颜色
}

type Keyword struct {
  Color string `json:"color"` // 当前文字颜色
  Value string `json:"value"` // 模板中占位符的值
}
