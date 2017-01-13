package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 单发模板消息
type AlipayOpenPublicMessageSingleSendRequest struct {
  api.IAlipayRequest
  TerminalType      string                                              `json:"terminal_type"`
  TerminalInfo      string                                              `json:"terminal_info"`
  ProdCode          string                                              `json:"prod_code"`
  NotifyUrl         string                                              `json:"notify_url"`
  ReturnUrl         string                                              `json:"return_url"`
  BizContent        AlipayOpenPublicMessageSingleSendRequestBizContent  `json:"biz_content"`
}

type AlipayOpenPublicMessageSingleSendRequestBizContent struct {
  ToUserId string `json:"to_user_id"`
  Template Template `json:"template"`
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetApiMethodName() string {
  return "alipay.open.public.message.single.send"
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicMessageSingleSendRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
