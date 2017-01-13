package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 群发消息
type AlipayOpenPublicMessageTotalSendRequest struct {
  api.IAlipayRequest
  TerminalType      string                                            `json:"terminal_type"`
  TerminalInfo      string                                            `json:"terminal_info"`
  ProdCode          string                                            `json:"prod_code"`
  NotifyUrl         string                                            `json:"notify_url"`
  ReturnUrl         string                                            `json:"return_url"`
  BizContent        AlipayOpenPublicMessageTotalSendRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicMessageTotalSendRequestBizContent struct {
  MsgType   string    `json:"msg_type"` // 消息类型，text：文本消息，image-text：图文消息
  Articles  []Article `json:"articles"` // 图文消息，当msg_type为image-text，该值必须设置
  Text      Text      `json:"text"`     // 文本消息内容，当msg_type为text，必须设置该值
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetApiMethodName() string {
  return "alipay.open.public.message.total.send"
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicMessageTotalSendRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicMessageTotalSendRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
