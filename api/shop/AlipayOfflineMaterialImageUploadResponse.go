package shop

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOfflineMaterialImageUploadResponse struct {
  api.AlipayResponse
  ImageId   string `json:"image_id"`    // 图片/视频在商家中心的唯一标识
  ImageURL  string `json:"image_url"`   // 图片/视频的访问地址（为了防止盗链，该地址不允许嵌在其他页面展示，只能在新页面展示）
}
