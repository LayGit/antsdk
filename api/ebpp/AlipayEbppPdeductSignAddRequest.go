package ebpp

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 缴费直连代扣签约
// 缴费直连代扣签约
type AlipayEbppPdeductSignAddRequest struct {
  api.IAlipayRequest
  TerminalType      string  `json:"terminal_type"`
  TerminalInfo      string  `json:"terminal_info"`
  ProdCode          string  `json:"prod_code"`
  NotifyUrl         string  `json:"notify_url"`
  ReturnUrl         string  `json:"return_url"`
  PayPasswordToken  string  `json:"pay_password_token"`   // 用户签约时，跳转到支付宝独立密码校验页面，校验成功后会将token和对应的用户ID缓存下来，然后跳回到机构页面生成token带回给机构，机构签约时必须传入token
  UserId            string  `json:"user_id"`              // 用户ID
  BizType           string  `json:"biz_type"`             // 业务类型。 JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用 WUYE：缴物业费 HK：信用卡还款 TX：手机充值
  SubBizType        string  `json:"sub_biz_type"`         // 业务子类型。 WATER：缴水费 ELECTRIC：缴电费 GAS：缴燃气费 COMMUN：缴固话宽带 CATV：缴有线电视费 TRAFFIC：缴交通罚款 WUYE：缴物业费 HK：信用卡还款 CZ：手机充值
  BillKey           string  `json:"bill_key"`             // 户号，机构针对于每户的水、电都会有唯一的标识户号
  OwnerName         string  `json:"owner_name"`           // 户名，户主真实姓名
  Pid               string  `json:"pid"`                  // 商户ID
  ChargeInst        string  `json:"charge_inst"`          // 支付宝缴费系统中的出账机构ID
  OutAgreementId    string  `json:"out_agreement_id"`     // 外部产生的协议ID
  AgentChannel      string  `json:"agent_channel"`        // 机构签约代扣来源渠道 PUBLICPLATFORM：服务窗
  AgentCode         string  `json:"agent_code"`           // 从服务窗发起则为publicId的值
  SignExpireDate    string  `json:"sign_expire_date"`     // 签约到期时间。空表示无限期，一期固定传空。
  PayConfig         string  `json:"pay_config"`           // 支付工具设置，目前可为空
  NotifyConfig      string  `json:"notify_config"`        // 通知方式设置，可为空
  DeductType        string  `json:"deduct_type"`          // 签约类型可为空
  ExtendField       string  `json:"extend_field"`         // 扩展字段
}

func (this *AlipayEbppPdeductSignAddRequest) GetApiMethodName() string {
  return "alipay.ebpp.pdeduct.sign.add"
}

func (this *AlipayEbppPdeductSignAddRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayEbppPdeductSignAddRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayEbppPdeductSignAddRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayEbppPdeductSignAddRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayEbppPdeductSignAddRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayEbppPdeductSignAddRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayEbppPdeductSignAddRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayEbppPdeductSignAddRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("pay_password_token", this.PayPasswordToken)
  txtParams.Put("user_id", this.UserId)
  txtParams.Put("biz_type", this.BizType)
  txtParams.Put("sub_biz_type", this.SubBizType)
  txtParams.Put("bill_key", this.BillKey)
  txtParams.Put("owner_name", this.OwnerName)
  txtParams.Put("pid", this.Pid)
  txtParams.Put("charge_inst", this.ChargeInst)
  txtParams.Put("out_agreement_id", this.OutAgreementId)
  txtParams.Put("agent_channel", this.AgentChannel)
  txtParams.Put("agent_code", this.AgentCode)
  txtParams.Put("sign_expire_date", this.SignExpireDate)
  txtParams.Put("pay_config", this.PayConfig)
  txtParams.Put("notify_config", this.NotifyConfig)
  txtParams.Put("deduct_type", this.DeductType)
  txtParams.Put("extend_field", this.ExtendField)
  return txtParams
}
