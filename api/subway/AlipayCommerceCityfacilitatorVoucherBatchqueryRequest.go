package subway

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 地铁购票订单批量查询
// 商户APP调该接口批量查询订单状态
type AlipayCommerceCityfacilitatorVoucherBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                          `json:"terminal_type"`
  TerminalInfo      string                                                          `json:"terminal_info"`
  ProdCode          string                                                          `json:"prod_code"`
  NotifyUrl         string                                                          `json:"notify_url"`
  ReturnUrl         string                                                          `json:"return_url"`
  BizContent        AlipayCommerceCityfacilitatorVoucherBatchqueryRequestBizContent `json:"biz_content"`
}

type AlipayCommerceCityfacilitatorVoucherBatchqueryRequestBizContent struct {
  CityCode string   `json:"city_code"`
  TradeNos []string `json:"trade_nos"`
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetApiMethodName() string {
  return "alipay.commerce.cityfacilitator.voucher.batchquery"
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayCommerceCityfacilitatorVoucherBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
