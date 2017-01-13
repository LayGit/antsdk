package open

type CodeInfo struct {
  Scene   Scene   `json:"scene"`      // 场景信息
  GotoUrl string  `json:"goto_url"`   // 跳转URL，扫码关注服务窗后会直接跳转到此URL
}

type Scene struct {
   SceneId string `json:"scene_id"` // 场景Id，最长32位，英文字母、数字以及下划线，开发者自定义
}
