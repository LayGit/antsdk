package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// isv 回传的用户操作行为信息调用接口
type AlipayOfflineProviderUseractionRecordRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                  `json:"terminal_type"`
  TerminalInfo        string                                                  `json:"terminal_info"`
  ProdCode            string                                                  `json:"prod_code"`
  NotifyUrl           string                                                  `json:"notify_url"`
  ReturnUrl           string                                                  `json:"return_url"`
  BizContent          AlipayOfflineProviderUseractionRecordRequestBizContent  `json:"biz_content"`
}

type AlipayOfflineProviderUseractionRecordRequestBizContent struct {
  ActionType      string      `json:"action_type"`
  Mobile          string      `json:"mobile"`
  PlatformUserId  string      `json:"platform_user_id"`
  Source          string      `json:"source"`
  Industry        string      `json:"industry"`
  UserId          string      `json:"user_id"`
  DateTime        string      `json:"date_time"`
  AlipayAppId     string      `json:"alipay_app_id"`
  ActionDetail    string      `json:"action_detail"`
  OuterShopDO     OuterShopDO `json:"outer_shop_do"`
  Entity          string      `json:"entity"`
  ActionOuterId   string      `json:"action_outer_id"`
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetApiMethodName() string {
  return "alipay.offline.provider.useraction.record"
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineProviderUseractionRecordRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineProviderUseractionRecordRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
