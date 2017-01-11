package coupon

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 支付宝pass新建卡券实例接口
// 卡券模板生成后，如需将卡券发布给对应的用户，则调用此接口。
type AlipayPassInstanceAddRequest struct {
  api.IAlipayRequest
  TerminalType      string                                   `json:"terminal_type"`
  TerminalInfo      string                                   `json:"terminal_info"`
  ProdCode          string                                   `json:"prod_code"`
  NotifyUrl         string                                   `json:"notify_url"`
  ReturnUrl         string                                   `json:"return_url"`
  BizContent        AlipayPassInstanceAddRequestBizContent   `json:"biz_content"`
}

type AlipayPassInstanceAddRequestBizContent struct {
  TplId           string `json:"tpl_id"`            // 支付宝pass模版ID
  TplParams       string `json:"tpl_params"`        // 模版动态参数信息【支付宝pass模版参数键值对JSON字符串】。
  RecognitionType string `json:"recognition_type"`  // Alipass添加对象识别类型：1–订单信息
  RecognitionInfo string `json:"recognition_info"`  // 支付宝用户识别信息： 包括partner_id（商户的签约账号）和out_trade_no
}

func (this *AlipayPassInstanceAddRequest) GetApiMethodName() string {
  return "alipay.pass.instance.add"
}

func (this *AlipayPassInstanceAddRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayPassInstanceAddRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayPassInstanceAddRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayPassInstanceAddRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayPassInstanceAddRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayPassInstanceAddRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayPassInstanceAddRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayPassInstanceAddRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
