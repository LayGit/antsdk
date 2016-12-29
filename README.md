# antsdk
蚂蚁金服(支付宝)开放平台 go-sdk

已支持下列接口：

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

描述 | API | Method
---|---|---
创建门店信息 | shop.AlipayOfflineMarketShopCreateRequest |  alipay.offline.market.shop.create
修改门店信息 | shop.AlipayOfflineMarketShopModifyRequest | alipay.offline.market.shop.modify
查询商户的门店编号列表 | shop.AlipayOfflineMarketShopBatchqueryRequest | alipay.offline.market.shop.batchquery
查询单个门店信息接口 | shop.AlipayOfflineMarketShopQuerydetailRequest | alipay.offline.market.shop.querydetail
上传门店照片和视频接口 | shop.AlipayOfflineMaterialImageUploadRequest | alipay.offline.material.image.upload
业务流水批量查询接口 | shop.AlipayOfflineMarketApplyorderBatchqueryRequest | alipay.offline.market.applyorder.batchquery
门店摘要信息批量查询接口 | shop.AlipayOfflineMarketShopSummaryBatchqueryRequest | alipay.offline.market.shop.summary.batchquery
门店类目配置查询接口 | shop.AlipayOfflineMarketShopCategoryQueryRequest | alipay.offline.market.shop.category.query

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
