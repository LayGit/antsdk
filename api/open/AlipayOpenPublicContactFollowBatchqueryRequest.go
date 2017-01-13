package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 查询服务窗联系人关注列表
type AlipayOpenPublicContactFollowBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                    `json:"terminal_type"`
  TerminalInfo      string                                                    `json:"terminal_info"`
  ProdCode          string                                                    `json:"prod_code"`
  NotifyUrl         string                                                    `json:"notify_url"`
  ReturnUrl         string                                                    `json:"return_url"`
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetApiMethodName() string {
  return "alipay.open.public.contact.follow.batchquery"
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicContactFollowBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
