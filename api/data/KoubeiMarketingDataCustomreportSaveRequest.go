package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 自定义报表规则创建及更新接口
type KoubeiMarketingDataCustomreportSaveRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                `json:"terminal_type"`
  TerminalInfo        string                                                `json:"terminal_info"`
  ProdCode            string                                                `json:"prod_code"`
  NotifyUrl           string                                                `json:"notify_url"`
  ReturnUrl           string                                                `json:"return_url"`
  BizContent          KoubeiMarketingDataCustomreportSaveRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportSaveRequestBizContent struct {
  ReportConditionInfo CustomReportCondition `json:"report_condition_info"`
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetApiMethodName() string {
  return "koubei.marketing.data.customreport.save"
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
