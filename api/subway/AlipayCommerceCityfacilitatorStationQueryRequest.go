package subway

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 地铁购票站点数据查询
// 商户App获取地铁购票开放购票的站点和禁止到达的站点
type AlipayCommerceCityfacilitatorStationQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                      `json:"terminal_type"`
  TerminalInfo      string                                                      `json:"terminal_info"`
  ProdCode          string                                                      `json:"prod_code"`
  NotifyUrl         string                                                      `json:"notify_url"`
  ReturnUrl         string                                                      `json:"return_url"`
  BizContent        AlipayCommerceCityfacilitatorStationQueryRequestBizContent  `json:"biz_content"`
}

type AlipayCommerceCityfacilitatorStationQueryRequestBizContent struct {
  CityCode string `json:"city_code"` // 城市编码请参考查询：http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201504/t20150415_712722.html； 已支持城市：广州 440100，深圳 440300，杭州330100。
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetApiMethodName() string {
  return "alipay.commerce.cityfacilitator.station.query"
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayCommerceCityfacilitatorStationQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
