package risk

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 蚁盾风险评分服务上数版
// “蚁盾”风险评分是综合利用大数据资产，以及层次化特征管理系统，对手机号等主体进行风险预测、风险解释的评分体系。产品已被众多商户应用于O2O、电商、团购、出行、手机3C等行业，针对批量机器注册、恶意刷单、黄牛抢购等风险场景，识别性能优异。
type SsdataDataserviceRiskRainscoreQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                  `json:"terminal_type"`
  TerminalInfo      string                                                  `json:"terminal_info"`
  ProdCode          string                                                  `json:"prod_code"`
  NotifyUrl         string                                                  `json:"notify_url"`
  ReturnUrl         string                                                  `json:"return_url"`
  BizContent        SsdataDataserviceRiskRainscoreQueryRequestBizContent    `json:"biz_content"`
}

type SsdataDataserviceRiskRainscoreQueryRequestBizContent struct {
  AccountType string `json:"account_type"`  // 账号类型，目前仅支持手机号（MOBILE_NO）
  Account     string `json:"account"`       // 帐号内容，目前为中国大陆手机号（11位阿拉伯数字，不包含特殊符号或空格）
  Version     string `json:"version"`       // “蚁盾”风险评分服务版本号，当前版本为2.0
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetApiMethodName() string {
  return "ssdata.dataservice.risk.rainscore.query"
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *SsdataDataserviceRiskRainscoreQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
