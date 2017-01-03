package zhima

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 获取数据反馈模板
// isv商户获取数据反馈模板
type ZhimaDataFeedbackurlQueryRequest struct {
  api.IAlipayRequest
  TerminalType  string                                      `json:"terminal_type"`
  TerminalInfo  string                                      `json:"terminal_info"`
  ProdCode      string                                      `json:"prod_code"`
  NotifyUrl     string                                      `json:"notify_url"`
  ReturnUrl     string                                      `json:"return_url"`
  BizContent    ZhimaDataFeedbackurlQueryRequestBizContent  `json:"biz_content"`
}

type ZhimaDataFeedbackurlQueryRequestBizContent struct {
  MerchantId string `json:"merchant_id"`
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetApiMethodName() string {
  return "zhima.data.feedbackurl.query"
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *ZhimaDataFeedbackurlQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
