package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 更新菜单
// 通过POST一个特定结构体，实现支付宝钱包客户端的公众账号更新自定义菜单。每一次的更新是针对全部自定义菜单的更新。
type AlipayOpenPublicMenuModifyRequest struct {
  api.IAlipayRequest
  TerminalType      string                                       `json:"terminal_type"`
  TerminalInfo      string                                       `json:"terminal_info"`
  ProdCode          string                                       `json:"prod_code"`
  NotifyUrl         string                                       `json:"notify_url"`
  ReturnUrl         string                                       `json:"return_url"`
  BizContent        AlipayOpenPublicMenuModifyRequestBizContent  `json:"biz_content"`
}

type AlipayOpenPublicMenuModifyRequestBizContent struct {
  Button []ButtonObject `json:"button"` // 一级菜单数组，个数应为1~4个
}

func (this *AlipayOpenPublicMenuModifyRequest) GetApiMethodName() string {
  return "alipay.open.public.menu.modify"
}

func (this *AlipayOpenPublicMenuModifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicMenuModifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicMenuModifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicMenuModifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicMenuModifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicMenuModifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicMenuModifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicMenuModifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
