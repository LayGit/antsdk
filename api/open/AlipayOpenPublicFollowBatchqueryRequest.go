package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 获取关注者列表
// 公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串userId组成。一次拉取调用最多拉取10000个关注者的userId，可以通过多次拉取的方式来满足需求。 公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串userId组成。一次拉取调用最多拉取10000个关注者的userId，可以通过多次拉取的方式来满足需求。
type AlipayOpenPublicFollowBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                            `json:"terminal_type"`
  TerminalInfo      string                                            `json:"terminal_info"`
  ProdCode          string                                            `json:"prod_code"`
  NotifyUrl         string                                            `json:"notify_url"`
  ReturnUrl         string                                            `json:"return_url"`
  BizContent        AlipayOpenPublicFollowBatchqueryRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicFollowBatchqueryRequestBizContent struct {
  NextUserId string `json:"next_user_id"`
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetApiMethodName() string {
  return "alipay.open.public.follow.batchquery"
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicFollowBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
