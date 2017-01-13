package marketing

type TemplateOpenCardConfDTO struct {
  OpenCardSourceType  string `json:"open_card_source_type"`  // ISV：外部系统 MER：直连商户
  SourceAppId         string `json:"source_app_id"`          // 渠道APPID，提供领卡页面的服务提供方
  OpenCardUrl         string `json:"open_card_url"`          // 开卡连接，必须http、https开头
  Conf                string `json:"conf"`                   // 配置，预留字段，暂时不用
}
