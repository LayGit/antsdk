package zhima

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 认证初始化
// 在使用认证接口之前(zhima.customer.certification.certify),需要先调用认证初始化接口,获取biz_no,然后在通过biz_no进行认证
type ZhimaCustomerCertificationInitializeRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                  `json:"terminal_type"`
  TerminalInfo      string                                                  `json:"terminal_info"`
  ProdCode          string                                                  `json:"prod_code"`
  NotifyUrl         string                                                  `json:"notify_url"`
  ReturnUrl         string                                                  `json:"return_url"`
  BizContent        ZhimaCustomerCertificationInitializeRequestBizContent   `json:"biz_content"`
}

type ZhimaCustomerCertificationInitializeRequestBizContent struct {
  TransactionId string `json:"transaction_id"`  // 商户请求的唯一标志，32位长度的字母数字下划线组合。该标识作为对账的关键信息，商户要保证其唯一性.建议:前面几位字符是商户自定义的简称,中间可以使用一段日期,结尾可以使用一个序列
  ProductCode   string `json:"product_code"`    // 芝麻认证产品码,示例值为真实的产品码
  BizCode       string `json:"biz_code"`        // 认证场景码,常用的场景码有: FACE:人脸认证, BANK_CARD:银行卡认证, ZFB_PASSWORD:支付宝帐号和支付密码密码认证 签约的协议决定了可以使用那些场景
  IdentityParam string `json:"identity_param"`  // 值为一个json串,必须包含身份类型identity_type,不同的身份类型需要的身份信息不同 当前支持: 身份信息为证件信息identity_type=CERT_INFO: 证件类型为身份证cert_type=IDENTITY_CARD,必要信息cert_name和cert_no 可以选填商户的用户主体principal_id,对应用户在商户端唯一标识,如果商户传了principal_id,后续会为商户提供更强大功能
  ExtBizParam   string `json:"ext_biz_param"`   // 扩展业务参数,暂时没有用到,接口预留
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetApiMethodName() string {
  return "zhima.customer.certification.initialize"
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetApiVersion() string {
  return "1.0"
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *ZhimaCustomerCertificationInitializeRequest) IsNeedEncrypt() bool {
  return false
}

func (this *ZhimaCustomerCertificationInitializeRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
