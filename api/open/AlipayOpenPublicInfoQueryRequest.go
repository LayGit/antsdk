package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务窗基础信息查询接口
// 通过该接口查询服务窗名称、头像、欢迎语基础信息
type AlipayOpenPublicInfoQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                       `json:"terminal_type"`
  TerminalInfo      string                                       `json:"terminal_info"`
  ProdCode          string                                       `json:"prod_code"`
  NotifyUrl         string                                       `json:"notify_url"`
  ReturnUrl         string                                       `json:"return_url"`
}

func (this *AlipayOpenPublicInfoQueryRequest) GetApiMethodName() string {
  return "alipay.open.public.info.query"
}

func (this *AlipayOpenPublicInfoQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicInfoQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicInfoQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicInfoQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicInfoQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicInfoQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicInfoQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicInfoQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  return txtParams
}
