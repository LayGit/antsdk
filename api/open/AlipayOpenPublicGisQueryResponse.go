package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicGisQueryResponse struct {
  api.AlipayResponse
  Latitude  string `json:"latitude"`    // 纬度信息
  Longitude string `json:"longitude"`   // 经度信息
  Accuracy  string `json:"accuracy"`    // 精确度
  Province  string `json:"province"`    // 经纬度对应位置所在的省份
  City      string `json:"city"`        // 经纬度所在位置
}
