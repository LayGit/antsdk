package coupon

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 支付宝pass更新卡券实例接口
// 对于已经发布的卡券，如需更新内容，可通过此接口更新，主要用于更新卡券的使用状态。
type AlipayPassInstanceUpdateRequest struct {
  api.IAlipayRequest
  TerminalType      string                                    `json:"terminal_type"`
  TerminalInfo      string                                    `json:"terminal_info"`
  ProdCode          string                                    `json:"prod_code"`
  NotifyUrl         string                                    `json:"notify_url"`
  ReturnUrl         string                                    `json:"return_url"`
  BizContent        AlipayPassInstanceUpdateRequestBizContent `json:"biz_content"`
}

type AlipayPassInstanceUpdateRequestBizContent struct {
  SerialNumber  string `json:"serial_number"`   // 商户指定卡券唯一值
  ChannelId     string `json:"channel_id"`      // 代理商代替商户发放卡券后，再代替商户更新卡券时，此值为商户的pid/appid
  TplParams     string `json:"tpl_params"`      // 模版动态参数信息【支付宝pass模版参数键值对JSON字符串】
  Status        string `json:"status"`          // 券状态，支持更新为USED、CLOSED两种状态
  VerifyCode    string `json:"verify_code"`     // 核销码串值【当状态变更为USED时，建议传】
  VerifyType    string `json:"verify_type"`     // 核销方式，目前支持：wave（声波方式）、qrcode（二维码方式）、barcode（条码方式）、input（文本方式，即手工输入方式）。pass和verify_type不能同时为空
}

func (this *AlipayPassInstanceUpdateRequest) GetApiMethodName() string {
  return "alipay.pass.instance.update"
}

func (this *AlipayPassInstanceUpdateRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayPassInstanceUpdateRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayPassInstanceUpdateRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayPassInstanceUpdateRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayPassInstanceUpdateRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayPassInstanceUpdateRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayPassInstanceUpdateRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayPassInstanceUpdateRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
