package marketing

type PointCard struct {
  Type      string `json:"type"`        // 工具类型，目前支持： 集点卡：POINT_CARD
  Name      string `json:"name"`        // 工具的名称
  Desc      string `json:"desc"`        // 工具的描述
  StartTime string `json:"start_time"`  // 工具的有效期的起始时间
  EndTime   string `json:"end_time"`    // 工具的有效期的结束时间（必须晚于活动的结束时间）
  Logo      string `json:"logo"`        // 工具的LOGO文件ID
}
