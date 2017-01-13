package data

type FilterTag struct {
  Code        string `json:"code"`          // 过滤条件的标签code
  FilterItems string `json:"filter_items"`  // 分组过滤条件 注意：这个是JSON数组，必须以[开头，以]结尾，[]外层不能加双引号"",正确案例[{"operate": "EQ","value": "1" }]，错误案例："[{"operate": "EQ","value": "1" }]"
}
