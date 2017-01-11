package ad

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 修改广告接口
// 开发者帮助线下商家修改广告内容，如修改的是线上的广告内容，需要先将线上广告内容下架，再修改，修改后操作上架，才能在支付宝钱包APP看到修改后的广告内容。运营位类型可以选择图片或H5。
type AlipayMarketingCdpAdvertiseModifyRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                `json:"terminal_type"`
  TerminalInfo      string                                                `json:"terminal_info"`
  ProdCode          string                                                `json:"prod_code"`
  NotifyUrl         string                                                `json:"notify_url"`
  ReturnUrl         string                                                `json:"return_url"`
  BizContent        AlipayMarketingCdpAdvertiseModifyRequestBizContent    `json:"biz_content"`
}

type AlipayMarketingCdpAdvertiseModifyRequestBizContent struct {
  AdId      string `json:"ad_id"`       // 广告ID,唯一标识一条广告
  Content   string `json:"content"`     // 广告内容。如果广告类型是HTML5，则传入H5链接地址，建议为https协议。最大尺寸不得超过1242px＊242px，小屏幕将按分辨率宽度同比例放大缩小；如果类型是图片，则传入图片ID标识，如何获取图片ID参考图片上传接口：alipay.offline.material.image.upload。图片尺寸为1242px＊290px。图片大小不能超过50kb。
  ActionUrl string `json:"action_url"`  // 行为地址。用户点击广告后，跳转URL地址, 协议必须为HTTPS。广告类型为PIC时，需要设置该值。对于类型为URL不生效
  Height    string `json:"height"`      // 当广告类型是H5时，必须传入内容高度。对于广告位CDP_OPEN_MERCHANT的内容高度不能高于钱包要求的展位高度242px。当广告类型是图片时，不需要传入内容高度(height)，系统会检查用户上传的图片尺寸是否符合要求，对于广告位CDP_OPEN_MERCHANT的图片尺寸要求：宽1242px, 高290px,大小50kb，实际上传图片与图片标准宽高必须一致，图片大小不能超过50kb。
  StartTime string `json:"start_time"`  // 投放广告开始时间，使用标准时间格式：yyyy-MM-dd HH:mm:ss，如果不设置，默认投放时间一个月
  EndTime   string `json:"end_time"`    // 投放广告结束时间，使用标准时间格式：yyyy-MM-dd HH:mm:ss，如果不设置，默认投放时间一个月
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetApiMethodName() string {
  return "alipay.marketing.cdp.advertise.modify"
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCdpAdvertiseModifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
