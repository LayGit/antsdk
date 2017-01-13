package data

type CustomReportCondition struct {
  ConditionKey  string      `json:"condition_key"`    // 规则KEY-为空则为创建规则，否则更新规则
  Name          string      `json:"name"`             // 自定义报表名称
  Memo          string      `json:"memo"`             // 规则描述
  DataTags      []DataTag   `json:"data_tags"`        // 明细数据标签
  GroupTags     string      `json:"group_tags"`       // 分组数据标签集合 注意：这个是JSON数组，必须以[开头，以]结尾，[]外层不能加双引号"",正确案例["orpt_ubase_age","orpt_ubase_birthday_mm"]，错误案例："["orpt_ubase_age","orpt_ubase_birthday_mm"]"
  SortTags      string      `json:"sort_tags"`        // 排序数据标签集合 注意：这个是JSON数组，必须以[开头，以]结尾，[]外层不能加双引号"",正确案例[{"code":"orpt_ubase_age","sortBy":"DESC"}]，错误案例："[{"code":"orpt_ubase_age","sortBy":"DESC"}]"
  FilterTags    []FilterTag `json:"filter_tags"`      // 分组过滤条件
}
