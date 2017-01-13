package marketing

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayMarketingCardTemplateQueryResponse struct {
  api.AlipayResponse
  CardType            string                    `json:"card_type"`              // 会员卡类型： OUT_MEMBER_CARD：外部权益卡
  BizNoPrefix         string                    `json:"biz_no_prefix"`          // 业务卡号前缀，由商户自定义
  BizNoSuffixLen      string                    `json:"biz_no_suffix_len"`      // 卡号长度
  TemplateStyleInfo   TemplateStyleInfoDTO      `json:"template_style_info"`    // 模板样式信息(钱包展现效果)
  TemplateBenefitInfo []TemplateBenefitInfoDTO  `json:"template_benefit_info"`  // 权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益
  ColumnInfoList      []TemplateColumnInfoDTO   `json:"column_info_list"`       // 栏位信息（卡包详情页面展现的栏位）
  FieldRuleList       TemplateFieldRuleDTO      `json:"field_rule_list"`        // 字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE
  ServiceLabelList    string                    `json:"service_label_list"`     // 服务标签列表
  ShopIds             []string                  `json:"shop_ids"`               // 门店ids
  OpenCardConf        TemplateOpenCardConfDTO   `json:"open_card_conf"`         // 会员卡用户领卡配置，在门店等渠道露出领卡入口时，需要部署的商户领卡H5页面地址
  PubChannels         []PubChannelDTO           `json:"pub_channels"`           // 卡模板投放渠道
  CardLevelConfs      []TemplateOpenCardConfDTO `json:"card_level_confs"`       // 卡等级配置
}
