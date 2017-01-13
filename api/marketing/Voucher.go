package marketing

type Voucher struct {
  Type                string              `json:"type"`                 // 券类型，目前支持以下类型： EXCHANGE：兑换券 MONEY：代金券 REDUCETO：减至券 RATE：折扣券
  VerifyMode          string              `json:"verify_mode"`          // 该字段仅在兑换券条件下(即券类型为EXCHANGE)，用于设置兑换券的核销方式 USER_CLICK:用户自己点击券上的按钮核销 MERCHANT_SCAN：商户通过APP扫码核销 其他情况下此字段为空
  VoucherNote         string              `json:"voucher_note"`         // 券的备注
  DescDetailList      []VoucherDescDetail `json:"desc_detail_list"`     // 券使用说明描述列表
  MergeVerifyConfig   MergeVerifyConfig   `json:"merge_verify_config"`  // 券核销叠加标识
  MultiUseMode        string              `json:"multi_use_mode"`       // 券叠加的属性，NO_MULTI:不可叠加;MULTI_USE_WITH_SINGLE:全场优惠和单品优惠的叠加；MULTI_USE_WITH_OTHERS:全场和其他所有优惠都可以叠加
  Name                string              `json:"name"`                 // 名称
  BrandName           string              `json:"brand_name"`           // 券副标题
  UseInstructions     []string            `json:"use_instructions"`     // 券的使用说明 使用须知最多6条，且每条最多100字
  Logo                string              `json:"logo"`                 // 券LOGO文件ID，调用图片上传接口alipay.offline.material.image.upload获得
  ItemInfo            ItemInfo            `json:"item_info"`            // 单品信息 兑换券不允许设置单品信息 减至券必须设置单品信息 其他类型券可按需设置
  BoucherImg          string              `json:"voucher_img"`          // 券图片文件ID,调用上传图片接口alipay.offline.material.image.upload获得
  Rate                string              `json:"rate"`                 // 折扣率 仅当券类型为折扣券时有效 有效折扣率取值范围0.11-0.99 仅允许保留小数点后两位
  MaxAmount           string              `json:"max_amount"`           // 最高优惠金额，单位元 必须为合法金额类型字符串 仅当券类型为DISOUNT有效
  WorthValue          string              `json:"worth_value"`          // 券面额，单位元 必须为合法金额类型字符串 券类型为代金券、减至券时，券面额必须设置 单品减至券的券面额必须低于单品原价
  DonateFlag          string              `json:"donate_flag"`          // 券是否可转赠，默认为可转赠
  ValidateType        string              `json:"validate_type"`        // 券有效期类型，目前支持以下类型： RELATIVE：相对有效期 FIXED：绝对有效期
  StartTime           string              `json:"start_time"`           // 券有效期的开始时间 仅在券有效期类型为绝对有效期时生效
  EndTime             string              `json:"end_time"`             // 券有效期的结束时间 仅在券有效期类型为绝对有效期时生效 必须晚于活动结束时间
  RelativeTime        string              `json:"relative_time"`        // 券相对有效期，单位天 仅在券有效期类型为相对有效期时生效 如，设5表示领券领取后5日内有效
  UseRule             UseRule             `json:"use_rule"`             // 券的使用规则信息
  Desc                string              `json:"desc"`                 // 券详细说明 最多包含500个字符
  DisplayConfig       DisplayConfig       `json:"display_config"`       // 券的展示信息
  ClauseTerms         []ClauseTerm        `json:"clause_terms"`         // 券的说明条款
  EffectType          string              `json:"effect_type"`          // 券生效的方式，目前支持以下方式 立即生效：IMMEDIATELY 延迟生效：DELAY 仅在券有效期类型为相对有效期时生效
  DelayInfo           DelayInfo           `json:"delay_info"`           // 延迟生效信息
  ExtInfo             string              `json:"ext_info"`             // 券的扩展信息
}
