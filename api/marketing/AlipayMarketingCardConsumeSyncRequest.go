package marketing

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 会员卡消费记录同步
type AlipayMarketingCardConsumeSyncRequest struct {
  api.IAlipayRequest
  TerminalType      string                                           `json:"terminal_type"`
  TerminalInfo      string                                           `json:"terminal_info"`
  ProdCode          string                                           `json:"prod_code"`
  NotifyUrl         string                                           `json:"notify_url"`
  ReturnUrl         string                                           `json:"return_url"`
  BizContent        AlipayMarketingCardConsumeSyncRequestBizContent  `json:"biz_content"`
}

type AlipayMarketingCardConsumeSyncRequestBizContent struct {
  TargetCardNo      string              `json:"target_card_no"`       // 支付宝业务卡号，开卡接口中返回获取
  TargetCardNoType  string              `json:"target_card_no_type"`  // 卡号类型 BIZ_CARD：支付宝业务卡号
  TradeType         string              `json:"trade_type"`           // 交易类型 开卡：OPEN 消费：TRADE 充值：DEPOSIT 退卡：RETURN
  TradeNo           string              `json:"trade_no"`             // 支付宝交易号
  TradeTime         string              `json:"trade_time"`           // 线下交易时间（是用户付款的交易时间） 当交易时间晚于上次消费记录同步时间，则会发生卡信息变更
  TradeName         string              `json:"trade_name"`           // 交易名称 为空时根据交易类型提供默认名称
  TradeAmount       float64             `json:"trade_amount"`         // 交易金额：本次交易的实际总金额（可认为标价金额）
  ActPayAmount      float64             `json:"act_pay_amount"`       // 用户实际付的现金金额 1.针对预付卡面额的核销金额在use_benefit_list展现，作为权益金额 2.权益金额不叠加在该金额上
  ShopCode          string              `json:"shop_code"`            // 门店编号
  Memo              string              `json:"memo"`                 // 备注信息，现有直接填写门店信息
  UseBenefitList    []BenefitInfoDetail `json:"use_benefit_list"`     // 实际消耗的权益
  GainBenefitList   []BenefitInfoDetail `json:"gain_benefit_list"`    // 获取权益列表
  CardInfo          MerchantCard        `json:"card_info"`            // 卡信息（交易后的实际卡信息）
  SwipeCertType     string              `json:"swipe_cert_type"`      // 产生该笔交易时，用户出具的凭证类型 ALIPAY：支付宝电子卡 ENTITY：实体卡 OTHER：其他
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetApiMethodName() string {
  return "alipay.marketing.card.consume.sync"
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayMarketingCardConsumeSyncRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayMarketingCardConsumeSyncRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
