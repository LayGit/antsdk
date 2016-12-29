package trade

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 统一收单交易支付
// 收银员使用扫码设备读取用户手机支付宝“付款码”/声波获取设备（如麦克风）读取用户手机支付宝的声波信息后，将二维码或条码信息/声波信息通过本接口上送至支付宝发起支付。
type AlipayTradePayRequest struct {
  api.IAlipayRequest
  TerminalType  string                            `json:"terminal_type"`
  TerminalInfo  string                            `json:"terminal_info"`
  ProdCode      string                            `json:"prod_code"`
  NotifyUrl     string                            `json:"notify_url"`
  ReturnUrl     string                            `json:"return_url"`
  BizContent    AlipayTradePayRequestBizContent   `json:"biz_content"`
}

type AlipayTradePayRequestBizContent struct {
  OutTradeNo            string          `json:"out_trade_no"`           // 商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
  Scene                 string          `json:"scene"`                  // 支付场景 条码支付，取值：bar_code 声波支付，取值：wave_code
  AuthCode              string          `json:"auth_code"`              // 支付授权码
  Subject               string          `json:"subject"`                // 订单标题
  SellerId              string          `json:"seller_id"`              // 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
  TotalAmount           float64         `json:"total_amount"`           // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
  DiscountableAmount    float64         `json:"discountable_amount"`    // 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
  UnDiscountableAmount  float64         `json:"undiscountable_amount"`  // 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【可打折金额】，则该值默认为【订单总金额】-【可打折金额】
  Body                  string          `json:"body"`                   // 订单描述
  GoodsDetail           []GoodsDetail   `json:"goods_detail"`           // 订单包含的商品列表信息，Json格式，其它说明详见商品明细说明
  OperatorId            string          `json:"operator_id"`            // 商户操作员编号
  StoreId               string          `json:"store_id"`               // 商户门店编号
  TerminalId            string          `json:"terminal_id"`            // 商户机具终端编号
  AlipayStoreId         string          `json:"alipay_store_id"`        // 支付宝的店铺编号
  ExtendParams          ExtendParams    `json:"extend_params"`          // 业务扩展参数
  TimeoutExpress        string          `json:"timeout_express"`        // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m
  RoyaltyInfo           RoyaltyInfo     `json:"royalty_info"`           // 描述分账信息，其它说明详见分账说明
  SubMerchant           SubMerchant     `json:"sub_merchant"`           // 二级商户信息,当前只对特殊银行机构特定场景下使用此字段
}

func (this *AlipayTradePayRequest) GetApiMethodName() string {
  return "alipay.trade.pay"
}

func (this *AlipayTradePayRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayTradePayRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayTradePayRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayTradePayRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayTradePayRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayTradePayRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayTradePayRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayTradePayRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
