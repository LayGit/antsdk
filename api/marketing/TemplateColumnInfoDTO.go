package marketing

type TemplateColumnInfoDTO struct {
  Code        string      `json:"code"`         // 标准栏位：行为由支付宝统一定，同时已经分配标准Code BALANCE：会员卡余额 POINT：积分 LEVEL：等级 TELEPHOME：联系方式 自定义栏位：行为由商户定义，自定义Code码（只要无重复）
  MoreInfo    MoreInfoDTO `json:"more_info"`    // 扩展信息
  Title       string      `json:"title"`        // 栏目的标题
  OperateType string      `json:"operate_type"` // 1、openNative：打开二级页面，展现 more中descs 2、openWeb：打开URL 3、staticinfo：静态信息
  Value       string      `json:"value"`        // 卡包详情页面，卡栏位右边展现的值
}

type MoreInfoDTO struct {
  Title   string   `json:"title"`   // 二级页面标题
  Url     string   `json:"url"`     // 超链接(选择openweb的时候必须填写url参数内容)
  Params  string   `json:"params"`  // 扩展参数，需要URL地址回带的值，JSON格式(openweb时填)
  Descs   []string `json:"descs"`   // 选择opennative的时候必须填写descs的内容
}
