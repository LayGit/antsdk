package coupon

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 发券授权
// 返回token授权商户给用户发奖
type KoubeiMarketingToolPrizesendAuthRequest struct {
  api.IAlipayRequest
  TerminalType      string                                              `json:"terminal_type"`
  TerminalInfo      string                                              `json:"terminal_info"`
  ProdCode          string                                              `json:"prod_code"`
  NotifyUrl         string                                              `json:"notify_url"`
  ReturnUrl         string                                              `json:"return_url"`
  BizContent        KoubeiMarketingToolPrizesendAuthRequestBizContent   `json:"biz_content"`
}

type KoubeiMarketingToolPrizesendAuthRequestBizContent struct {
  ReqId   string `json:"req_id"`    // 外部流水号，保证业务幂等性
  PrizeId string `json:"prize_id"`  // 奖品ID
  UserId  string `json:"user_id"`   // 发奖用户ID
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetApiMethodName() string {
  return "koubei.marketing.tool.prizesend.auth"
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
