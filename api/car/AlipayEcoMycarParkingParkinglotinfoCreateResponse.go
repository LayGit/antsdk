package car

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayEcoMycarParkingParkinglotinfoCreateResponse struct {
  api.AlipayResponse
  ParkingId string `json:"parking_id"`  // 支付宝返回停车场id。成功不为空，失败返回空
}
