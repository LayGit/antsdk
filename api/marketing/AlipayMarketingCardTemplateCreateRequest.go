package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 会员卡模板创建
type AlipayMarketingCardTemplateCreateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                             `json:"terminal_type"`
  TerminalInfo      string                                             `json:"terminal_info"`
  ProdCode          string                                             `json:"prod_code"`
  NotifyUrl         string                                             `json:"notify_url"`
  ReturnUrl         string                                             `json:"return_url"`
  BizContent        AlipayMarketingCardTemplateCreateRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCardTemplateCreateRequestBizContent struct {
  RequestId           string                      `json:"request_id"`             // 请求ID，由开发者生成并保证唯一性
  CardType            string                      `json:"card_type"`              // 卡类型为固定枚举类型，可选类型如下： OUT_MEMBER_CARD：外部权益卡
  BizNoPrefix         string                      `json:"biz_no_prefix"`          // 业务卡号前缀，由商户自定义
  BizNoSuffixLen      string                      `json:"biz_no_suffix_len"`      // 业务卡号后缀长度，最大int值不能超过32-biz_no_suffix长度
  WriteOffType        string                      `json:"write_off_type"`         // 卡包详情页面中展现出的卡码（可用于扫码核销） qrcode: 二维码 dqrcode: 动态二维码 barcode: 条码 dbarcode: 动态条码 text: 文本
  TemplateStyleInfo   TemplateStyleInfoDTO        `json:"template_style_info"`    // 模板样式信息
  TemplateBenefitInfo []TemplateBenefitInfoDTO    `json:"template_benefit_info"`  // 权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益
  ColumnInfoList      []TemplateColumnInfoDTO     `json:"column_info_list"`       // 栏位信息
  FieldRuleList       []TemplateFieldRuleDTO      `json:"field_rule_list"`        // 字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE
  OpenCardConf        TemplateOpenCardConfDTO     `json:"open_card_conf"`         // 会员卡用户领卡配置，在门店等渠道露出领卡入口时，需要部署的商户领卡H5页面地址
  ServiceLabelList    []string                    `json:"service_label_list"`     // 服务Code HUABEI_FUWU：花呗服务（只有需要花呗服务时，才需要加入该标识）
  ShopIds             []string                    `json:"shop_ids"`               // 会员卡上架门店id（支付宝门店id），既发放会员卡的商家门店id
  PubChannels         []PubChannelDTO             `json:"pub_channels"`           // 卡模板投放渠道
  CardLevelConf       []TemplateCardLevelConfDTO  `json:"card_level_conf"`        // 卡级别配置
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetApiMethodName() string {
  return "alipay.marketing.card.template.create"
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCardTemplateCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCardTemplateCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
