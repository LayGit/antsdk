package data

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// isv 回传的商户操作行为信息调用接口
type AlipayOfflineProviderShopactionRecordRequest struct {
  api.IAlipayRequest
  TerminalType        string                                                  `json:"terminal_type"`
  TerminalInfo        string                                                  `json:"terminal_info"`
  ProdCode            string                                                  `json:"prod_code"`
  NotifyUrl           string                                                  `json:"notify_url"`
  ReturnUrl           string                                                  `json:"return_url"`
  BizContent          AlipayOfflineProviderShopactionRecordRequestBizContent  `json:"biz_content"`
}

type AlipayOfflineProviderShopactionRecordRequestBizContent struct {
  ActionType    string      `json:"action_type"`      // 操作类型（insert_table/update_table/insert_dish/delete_dish/soldout_dish/modify_dish/modify_shop_status）
  Entity        string      `json:"entity"`           // 操作实体（实体+操作类型决定一个真正的操作【店铺+新增、座位+修改、店铺+适时状态等等】）
  DateTime      string      `json:"date_time"`        // 商户行为发生时间 格式：yyyy-MM-dd HH:mm:ss
  OuterShopDO   OuterShopDO `json:"outer_shop_do"`    // 操作的店铺对象
  ActionDetail  string      `json:"action_detail"`    // json格式，操作详情（根据操作类型不同而不同）（太大的话可能会导致内部处理失败）
  Industry      string      `json:"industry"`         // 所属行业 (餐饮：REPAST)
  ActionOuterId string      `json:"action_outer_id"`  // 本次请求的唯一键（操作实体主键+平台字符串）
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetApiMethodName() string {
  return "alipay.offline.provider.shopaction.record"
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineProviderShopactionRecordRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineProviderShopactionRecordRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
