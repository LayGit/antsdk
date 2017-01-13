package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 自定义报表规则列表分页查询接口
// 分页查询自定义数据报表规则列表
type KoubeiMarketingDataCustomreportBatchqueryRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                      `json:"terminal_type"`
  TerminalInfo        string                                                      `json:"terminal_info"`
  ProdCode            string                                                      `json:"prod_code"`
  NotifyUrl           string                                                      `json:"notify_url"`
  ReturnUrl           string                                                      `json:"return_url"`
  BizContent          KoubeiMarketingDataCustomreportBatchqueryRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportBatchqueryRequestBizContent struct {
  PageNo    string `json:"page_no"`   // 当前页号，默认为1
  PageSize  string `json:"page_size"` // 每页条目数，默认为20,最大为30
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetApiMethodName() string {
  return "koubei.marketing.data.customreport.batchquery"
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
