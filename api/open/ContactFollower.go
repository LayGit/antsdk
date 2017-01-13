package open

type ContactFollower struct {
  Avatar          string `json:"avatar"`            // 支付宝头像
  DefaultAvatar   string `json:"default_avatar"`    // 默认头像
  EachRecordFlag  string `json:"each_record_flag"`  // false
  UserId          string `json:"user_id"`           // 用户id
}
