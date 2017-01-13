package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 异步单发消息
type AlipayOpenPublicMessageCustomSendRequest struct {
  api.IAlipayRequest
  TerminalType      string                                               `json:"terminal_type"`
  TerminalInfo      string                                               `json:"terminal_info"`
  ProdCode          string                                               `json:"prod_code"`
  NotifyUrl         string                                               `json:"notify_url"`
  ReturnUrl         string                                               `json:"return_url"`
  BizContent        AlipayOpenPublicMessageCustomSendRequestBizContent   `json:"biz_content"`
}

type AlipayOpenPublicMessageCustomSendRequestBizContent struct {
  ToUserId  string    `json:"to_user_id"` // 消息接收用户的userid
  MsgType   string    `json:"msg_type"`   // 消息类型，text：文本消息，image-text：图文消息
  Articles  []Article `json:"articles"`   // 图文消息，当msg_type为image-text时，必须存在相对应的值
  Text      Text      `json:"text"`       // 当msg_type为text时，必须设置相对应的值
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetApiMethodName() string {
  return "alipay.open.public.message.custom.send"
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicMessageCustomSendRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicMessageCustomSendRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
