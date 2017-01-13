package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 口碑商户经营数据查询接口
type KoubeiMarketingDataIndicatorQueryRequest struct {
  api.IAlipayRequest
  TerminalType        string                                              `json:"terminal_type"`
  TerminalInfo        string                                              `json:"terminal_info"`
  ProdCode            string                                              `json:"prod_code"`
  NotifyUrl           string                                              `json:"notify_url"`
  ReturnUrl           string                                              `json:"return_url"`
  BizContent          KoubeiMarketingDataIndicatorQueryRequestBizContent  `json:"biz_content"`
}

type KoubeiMarketingDataIndicatorQueryRequestBizContent struct {
  BeginDate string `json:"begin_date"`  // 开始日期,格式:yyyyMMdd
  EndDate   string `json:"end_date"`    // 结束日期 格式:yyyyMMdd
  PageNum   string `json:"page_num"`    // 当前页数，默认为1
  PageSize  string `json:"page_size"`   // 每页记录数,不能超过50，默认为20
  BizType   string `json:"biz_type"`    // 业务类型，可选值有六个 1，MemberQuery 商户会员数据查询 2，MemberQueryByStore 门店会员数据查询 3，TradeQuery 商户交易数据查询 4，TradeQueryByStore 门店交易数据查询 5，CampaignQuery 商户活动数据查询 6，CampaignQueryByStore 门店活动数据查询
  ExtInfo   string `json:"ext_info"`    // camp_id：为活动ID sort_field：为排序指标KEY sort_type：ASC表示升序,DESC表示降序 store_Ids：为门店ID，多个门店使用逗号分隔
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetApiMethodName() string {
  return "koubei.marketing.data.indicator.query"
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingDataIndicatorQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
