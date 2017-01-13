package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 招商门店分页查询接口
// 活动创建时，在购物中心和品牌商场景需要对活动对应的门店进行招商，商家确认后，购物中心PID和品牌商登陆后，需要查下看招商情况
type KoubeiMarketingCampaignRecruitShopQueryRequest struct {
  api.IAlipayRequest
  TerminalType      string                                                   `json:"terminal_type"`
  TerminalInfo      string                                                   `json:"terminal_info"`
  ProdCode          string                                                   `json:"prod_code"`
  NotifyUrl         string                                                   `json:"notify_url"`
  ReturnUrl         string                                                   `json:"return_url"`
  BizContent        KoubeiMarketingCampaignRecruitShopQueryRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingCampaignRecruitShopQueryRequestBizContent struct {
  CampId        string  `json:"camp_id"`        // 活动id
  PageSize      int     `json:"page_size"`      // 每页大小
  PageNum       int     `json:"page_num"`       // 页码
  Invitee       string  `json:"invitee"`        // 参与的商户Id
  OperatorId    string  `json:"operator_id"`    // 操作人id
  OperatorType  string  `json:"operator_type"`  // 操作人类型
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetApiMethodName() string {
  return "koubei.marketing.campaign.recruit.shop.query"
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetApiVersion() string {
  return "1.0"
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) IsNeedEncrypt() bool {
  return false
}

func (this *KoubeiMarketingCampaignRecruitShopQueryRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
