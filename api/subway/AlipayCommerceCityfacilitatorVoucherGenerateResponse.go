package subway

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayCommerceCityfacilitatorVoucherGenerateResponse struct {
  api.AlipayResponse
  QRCodeNO    string `json:"qr_code_no"`    // 地铁购票二维码编码，可自定义
  TicketNo    string `json:"ticket_no"`     // 地铁购票的核销码
  ExpiredDate string `json:"expired_date"`  // 核销码过期时间
}
