package risk

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// “蚁盾”风险评分服务
// “蚁盾”风险评分是综合利用大数据资产，以及层次化特征管理系统，对手机号等主体进行风险预测、风险解释的评分体系。产品已被众多商户应用于O2O、电商、团购、出行、手机3C等行业，针对批量机器注册、恶意刷单、黄牛抢购等风险场景，识别性能优异。
type AlipaySecurityRiskRainscoreQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                               `json:"terminal_type"`
  TerminalInfo      string                                               `json:"terminal_info"`
  ProdCode          string                                               `json:"prod_code"`
  NotifyUrl         string                                               `json:"notify_url"`
  ReturnUrl         string                                               `json:"return_url"`
  BizContent        AlipaySecurityRiskRainscoreQueryRequestBizContent    `json:"biz_content"`
}

type AlipaySecurityRiskRainscoreQueryRequestBizContent struct {
  AccountType string `json:"account_type"`  // 账号类型，目前仅支持手机号（MOBILE_NO）
  Account     string `json:"account"`       // 帐号内容，目前为中国大陆手机号（11位阿拉伯数字，不包含特殊符号或空格）
  Version     string `json:"version"`       // “蚁盾”风险评分服务版本号，当前版本为2.0
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetApiMethodName() string {
  return "alipay.security.risk.rainscore.query"
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipaySecurityRiskRainscoreQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
