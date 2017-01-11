package finance

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 余利宝申购
type MybankFinanceYulibaoCapitalPurchaseRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                `json:"terminal_type"`
  TerminalInfo      string                                                `json:"terminal_info"`
  ProdCode          string                                                `json:"prod_code"`
  NotifyUrl         string                                                `json:"notify_url"`
  ReturnUrl         string                                                `json:"return_url"`
  BizContent        MybankFinanceYulibaoCapitalPurchaseRequestBizContent  `json:"biz_content"`
}

type MybankFinanceYulibaoCapitalPurchaseRequestBizContent struct {
  FundCode  string `json:"fund_code"`   // 基金代码，必填。目前默认填001529，代表余利宝
  Amount    int    `json:"amount"`      // 余利宝申购金额，单位是“分”。如amount=123456表示申购1234.56元的余利宝份额。
  Currency  string `json:"currency"`    // 金额对应的币种，目前仅支持CNY，即人民币。
  OutBizNo  string `json:"out_biz_no"`  // 余利宝申购流水号，用于幂等控制。流水号必须长度在30到40位之间，且仅能由数字、字母、字符“-”和字符“_”组成。建议使用UUID，如“c39c24f1-73e5-497d-9b1f-0f585ae192c1”，或者使用自定义的数字流水号，如“201608150000000000000000000000000001”。
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetApiMethodName() string {
  return "mybank.finance.yulibao.capital.purchase"
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetApiVersion() string {
  return "1.0"
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) IsNeedEncrypt() bool {
  return false
}

func (this *MybankFinanceYulibaoCapitalPurchaseRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
