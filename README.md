# antsdk
蚂蚁金服(支付宝)开放平台 go-sdk

## 使用示例

```go
import (
  "fmt"
  "github.com/LayGit/antsdk/alipay"
  "github.com/LayGit/antsdk/api/trade"
  "github.com/LayGit/antsdk/utils"
)

func main() {
    client := alipay.NewDefaultAlipayClient("https://openapi.alipay.com/gateway.do", "商户AppId", "商户密钥", "支付宝公钥")
    // 创建请求
    request := &trade.AlipayTradeQueryRequest{}
    // 设置参数
    request.BizContent.OutTradeNo = "L123456"
    // 请求响应
    var response trade.AlipayTradeQueryResponse
    err := client.Execute(request, &response)
    if err != nil {
        // 错误处理
        fmt.Println(err)
    } else {
        // todo:
        fmt.Println(response.Code)
    }
}
```

## API 支持情况
- [x] [支付 API](#支付-api)
- [x] [会员 API](#会员-api)
- [x] [店铺 API](#店铺-api)
- [ ] 营销 API
- [ ] 服务窗 API
- [x] [芝麻信用 API](#芝麻信用-api)
- [x] [工具类 API](#工具类-api)
- [ ] 风险控制 API
- [ ] 服务市场 API
- [ ] 账务 API
- [ ] 生活缴费 API
- [ ] 车主服务 API
- [ ] 数据服务 API
- [ ] 卡券 API
- [ ] 广告 API
- [ ] 地铁购票 API
- [ ] 理财 API

## 支付 API

描述 | API | Method
---|---|---
统一收单交易撤销接口 | trade.AlipayTradeCancelRequest | alipay.trade.cancel
统一收单交易关闭接口 | trade.AlipayTradeCloseRequest | alipay.trade.close
统一收单交易创建  | trade.AlipayTradeCreateRequest | alipay.trade.create
统一收单交易退款查询 | trade.AlipayTradeFastpayRefundQueryRequest | alipay.trade.fastpay.refund.query
统一收单交易结算接口 | trade.AlipayTradeOrderSettleRequest | alipay.trade.order.settle
统一收单交易支付 | trade.AlipayTradePayRequest | alipay.trade.pay
统一收单线下交易预创建 | trade.AlipayTradePreCreateRequest | alipay.trade.precreate
统一收单线下交易查询 | trade.AlipayTradeQueryRequest | alipay.trade.query
统一收单交易退款 | trade.AlipayTradeRefundRequest | alipay.trade.refund

## 会员 API

描述 | API | Method
---|---|---
支付宝钱包用户信息共享 | user.AlipayUserUserinfoShareRequest | alipay.user.userinfo.share
支付宝会员授权信息查询接口 | user.AlipayUserInfoShareRequest | alipay.user.info.share

## 店铺 API

描述 | API
---|---
创建门店信息 | shop.AlipayOfflineMarketShopCreateRequest
修改门店信息 | shop.AlipayOfflineMarketShopModifyRequest
查询商户的门店编号列表 | shop.AlipayOfflineMarketShopBatchqueryRequest
查询单个门店信息接口 | shop.AlipayOfflineMarketShopQuerydetailRequest
上传门店照片和视频接口 | shop.AlipayOfflineMaterialImageUploadRequest
业务流水批量查询接口 | shop.AlipayOfflineMarketApplyorderBatchqueryRequest
门店摘要信息批量查询接口 | shop.AlipayOfflineMarketShopSummaryBatchqueryRequest
门店类目配置查询接口 | shop.AlipayOfflineMarketShopCategoryQueryRequest

## 芝麻信用 API

描述 | API | Method
---|---|---
获取数据反馈模板 | zhima.ZhimaDataFeedbackurlQueryRequest | zhima.data.feedbackurl.query
批量数据反馈服务 | zhima.ZhimaDataBatchFeedbackRequest | zhima.data.batch.feedback
芝麻信用评分普惠版 | zhima.ZhimaCreditScoreBriefGetRequest | zhima.credit.score.brief.get
行业关注名单普惠版 | zhima.ZhimaCreditWatchlistBriefGetRequest | zhima.credit.watchlist.brief.get
欺诈信息验证 | zhima.ZhimaCreditAntifraudVerifyRequest | zhima.credit.antifraud.verify
认证初始化 | zhima.ZhimaCustomerCertificationInitializeRequest | zhima.customer.certification.initialize
芝麻认证开始认证 | zhima.ZhimaCustomerCertificationCertifyRequest | zhima.customer.certification.certify
芝麻认证查询 | zhima.ZhimaCustomerCertificationQueryRequest | zhima.customer.certification.query

## 工具类 API

描述 | API | Method
---|---|---
换取应用授权令牌 | util.AlipayOpenAuthTokenAppRequest | alipay.open.auth.token.app
换取授权访问令牌 | util.AlipaySystemOauthTokenRequest | alipay.system.oauth.token
查询某个应用授权AppAuthToken的授权信息 | util.AlipayOpenAuthTokenAppQueryRequest | alipay.open.auth.token.app.query
用户登陆授权 | util.AlipayUserInfoAuthRequest | alipay.user.info.auth
