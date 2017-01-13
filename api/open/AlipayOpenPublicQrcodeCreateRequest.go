package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 带参推广二维码接口
// 为了满足用户渠道推广分析的需要，公众平台提供了生成带参数二维码的接口。使用该接口可以获得多个带不同场景值的二维码，用户扫描后，公众号可以接收到事件推送。目前有2种类型的二维码，分别是临时二维码、和永久二维码，前者有过期时间，最大为1800秒。 每次创建二维码ticket需要提供一个开发者自行设定的参数（scene_id）。
type AlipayOpenPublicQrcodeCreateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                          `json:"terminal_type"`
  TerminalInfo      string                                          `json:"terminal_info"`
  ProdCode          string                                          `json:"prod_code"`
  NotifyUrl         string                                          `json:"notify_url"`
  ReturnUrl         string                                          `json:"return_url"`
  BizContent        AlipayOpenPublicQrcodeCreateRequestBizContent   `json:"biz_content"`
}

type AlipayOpenPublicQrcodeCreateRequestBizContent struct {
  CodeInfo      CodeInfo  `json:"code_info"`      // 服务窗创建带参二维码接口，开发者自定义信息
  CodeType      string    `json:"code_type"`      // 二维码类型，目前只支持两种类型： TEMP：临时的（默认）； PERM：永久的
  ExpireSecond  string    `json:"expire_second"`  // 临时码过期时间，以秒为单位，最大不超过1800秒； 永久码置空
  ShowLogo      string    `json:"show_logo"`      // 二维码中间是否显示服务窗logo，Y：显示；N：不显示（默认）
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetApiMethodName() string {
  return "alipay.open.public.qrcode.create"
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicQrcodeCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicQrcodeCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
