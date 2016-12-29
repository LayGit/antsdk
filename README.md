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
