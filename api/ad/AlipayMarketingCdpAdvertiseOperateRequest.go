package ad

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 操作广告接口
// 开发者帮助线下商家上线/下线广告内容，上线后用户可以在钱包APP上看到广告，下线后用户不能看到广告且该广告内容失效。
type AlipayMarketingCdpAdvertiseOperateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                `json:"terminal_type"`
  TerminalInfo      string                                                `json:"terminal_info"`
  ProdCode          string                                                `json:"prod_code"`
  NotifyUrl         string                                                `json:"notify_url"`
  ReturnUrl         string                                                `json:"return_url"`
  BizContent        AlipayMarketingCdpAdvertiseOperateRequestBizContent   `json:"biz_content"`
}

type AlipayMarketingCdpAdvertiseOperateRequestBizContent struct {
  AdId        string `json:"ad_id"`         // 广告ID，唯一标识一条广告
  OperateType string `json:"operate_type"`  // 操作类型，目前包括上线和下线，分别传入：online(ONLINE)和offline(OFFLINE)
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetApiMethodName() string {
  return "alipay.marketing.cdp.advertise.operate"
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCdpAdvertiseOperateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
