package zhima

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 芝麻认证查询
// 商户在认证完成后,可以查询认证的状态和结果
type ZhimaCustomerCertificationQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                              `json:"terminal_type"`
  TerminalInfo      string                                              `json:"terminal_info"`
  ProdCode          string                                              `json:"prod_code"`
  NotifyUrl         string                                              `json:"notify_url"`
  ReturnUrl         string                                              `json:"return_url"`
  BizContent        ZhimaCustomerCertificationQueryRequestBizContent    `json:"biz_content"`
}

type ZhimaCustomerCertificationQueryRequestBizContent struct {
  BizNo string `json:"biz_no"` // 一次认证的唯一标识,在商户调用认证初始化接口的时候获取
}

func (this *ZhimaCustomerCertificationQueryRequest) GetApiMethodName() string {
  return "zhima.customer.certification.query"
}

func (this *ZhimaCustomerCertificationQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *ZhimaCustomerCertificationQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *ZhimaCustomerCertificationQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *ZhimaCustomerCertificationQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *ZhimaCustomerCertificationQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *ZhimaCustomerCertificationQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *ZhimaCustomerCertificationQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *ZhimaCustomerCertificationQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
