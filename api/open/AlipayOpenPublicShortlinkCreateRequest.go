package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务窗短链自主生成接口
// 商户通过本接口生成带自有场景标识的短链接
type AlipayOpenPublicShortlinkCreateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                            `json:"terminal_type"`
  TerminalInfo      string                                            `json:"terminal_info"`
  ProdCode          string                                            `json:"prod_code"`
  NotifyUrl         string                                            `json:"notify_url"`
  ReturnUrl         string                                            `json:"return_url"`
  BizContent        AlipayOpenPublicShortlinkCreateRequestBizContent  `json:"biz_content"`
}

type AlipayOpenPublicShortlinkCreateRequestBizContent struct {
  SceneId string `json:"scene_id"`  // 短链接对应的场景ID，该ID由商户自己定义
  Remark  string `json:"remark"`    // 对于场景ID的描述，商户自己定义
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetApiMethodName() string {
  return "alipay.open.public.shortlink.create"
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicShortlinkCreateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
