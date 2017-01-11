package ad

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayMarketingCdpAdvertiseQueryResponse struct {
  api.AlipayResponse
  AdCode      string `json:"ad_code"`       // 广告位标识码，目前开放的广告位是钱包APP/口碑TAB/商家详情页中，传值：CDP_OPEN_MERCHANT
  ContentType string `json:"content_type"`  // 广告内容类型，目前包括HTML5和图片，分别传入：URL和PIC
  Content     string `json:"content"`       // 广告内容。如果广告类型是HTML5，则传入H5链接地址，建议为https协议。最大尺寸不得超过1242px＊242px，小屏幕将按分辨率宽度同比例放大缩小；如果类型是图片，则传入图片ID标识，如何获取图片ID参考图片上传接口：alipay.offline.material.image.upload。图片尺寸为1242px＊290px。图片大小不能超过50kb。
  ActionUrl   string `json:"action_url"`    // 用户点击广告后，跳转URL地址, 协议必须为HTTPS。广告类型为PIC时，需要设置该值。对于类型为URL不生效
  AdRules     string `json:"ad_rules"`      // 该规则用于商家设置对用户是否展示广告的校验条件，目前支持设置商家店铺规则。按业务要求应用对应规则即可
  Height      string `json:"height"`        // 当广告类型是H5时，必须传入内容高度。对于广告位CDP_OPEN_MERCHANT的内容高度不能高于钱包要求的展位高度242px。当广告类型是图片时，不需要传入内容高度(height)，系统会检查用户上传的图片尺寸是否符合要求，对于广告位CDP_OPEN_MERCHANT的图片尺寸要求：宽1242px, 高290px,大小50kb，实际上传图片与图片标准宽高必须一致，图片大小不能超过50kb。
  Status      string `json:"status"`        // 在线：ONLINE , 下线：OFFLINE
  StartTime   string `json:"start_time"`    // 投放广告开始时间，使用标准时间格式：yyyy-MM-dd HH:mm:ss，如果不设置，默认投放时间一个月
  EndTime     string `json:"end_time"`      // 投放广告结束时间，使用标准时间格式：yyyy-MM-dd HH:mm:ss，如果不设置，默认投放时间一个月
}
