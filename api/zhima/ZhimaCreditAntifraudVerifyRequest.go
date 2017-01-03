package zhima

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 欺诈信息验证
// 信息真实性和可信度验证，输出风险因素code
type ZhimaCreditAntifraudVerifyRequest struct {
  api.IAlipayRequest
  TerminalType      string                                        `json:"terminal_type"`
  TerminalInfo      string                                        `json:"terminal_info"`
  ProdCode          string                                        `json:"prod_code"`
  NotifyUrl         string                                        `json:"notify_url"`
  ReturnUrl         string                                        `json:"return_url"`
  BizContent        ZhimaCreditAntifraudVerifyRequestBizContent   `json:"biz_content"`
}

type ZhimaCreditAntifraudVerifyRequestBizContent struct {
  ProductCode   string `json:"product_code"`      // 产品码，标记商户接入的具体产品；直接使用每个接口入参中示例值即可
  TransactionId string `json:"transaction_id"`    // 商户请求的唯一标志，长度64位以内字符串，仅限字母数字下划线组合。该标识作为业务调用的唯一标识，商户要保证其业务唯一性，使用相同transaction_id的查询，芝麻在一段时间内（一般为1天）返回首次查询结果，超过有效期的查询即为无效并返回异常，有效期内的重复查询不重新计费
  CertNo        string `json:"cert_no"`           // 证件号
  CertType      string `json:"cert_type"`         // 证件类型。IDENTITY_CARD标识为身份证，目前仅支持身份证类型
  Name          string `json:"name"`              // 姓名，长度不超过64，姓名中不含",","/u0001"，"|","&","^","\\"
  Mobile        string `json:"mobile"`            // 手机号码。中国大陆合法手机号，长度11位，不含国家代码
  Email         string `json:"email"`             // 电子邮箱。合法email，字母小写，特殊符号以半角形式出现
  BankCard      string `json:"bank_card"`         // 银行卡号。中国大陆银行发布的银行卡:借记卡长度19位；信用卡长度16位；各位的取值位[0,9]的整数；不含虚拟卡
  Address       string `json:"address"`           // 地址信息。省+市+区/县+详细地址，其中 省+市+区/县可以为空，长度不超过256，不含",","/u0001"，"|","&","^","\\"
  IP            string `json:"ip"`                // ip地址。以"."分割的四段Ip，如 x.x.x.x，x为[0,255]之间的整数
  Mac           string `json:"mac"`               // 物理地址。支持格式如下：xx:xx:xx:xx:xx:xx，xx-xx-xx-xx-xx-xx，xxxxxxxxxxxx，x取值范围[0,9]之间的整数及A，B，C，D，E，F
  WiFiMac       string `json:"wifimac"`           // wifi的物理地址。支持格式如下：xx:xx:xx:xx:xx:xx，xx-xx-xx-xx-xx-xx，xxxxxxxxxxxx，x取值范围[0,9]之间的整数及A，B，C，D，E，F
  IMEI          string `json:"imei"`              // 国际移动设备标志。15位长度数字
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetApiMethodName() string {
  return "zhima.credit.antifraud.verify"
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *ZhimaCreditAntifraudVerifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *ZhimaCreditAntifraudVerifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
